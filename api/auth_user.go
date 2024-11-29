package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/go-http-server/core/utils"
	"github.com/go-http-server/core/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type RegisterUserRequestParams struct {
	Username string `json:"username" binding:"required,min=6,alphanum,lowercase"`
	Email    string `json:"email" binding:"required,min=12,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$?&*"`
	FullName string `json:"full_name" binding:"required,min=6"`
}

type LoginUserRequestParams struct {
	Identifier string `json:"identifier" binding:"required,min=6"`
	Password   string `json:"password" binding:"required,min=6,containsany=!@#$?&*"`
}

type UserResponseLoginRequest struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	FullName          string    `json:"full_name"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	IsVerifiedEmail   bool      `json:"is_verified_email"`
}

type loginResponseBody struct {
	AccessToken string                   `json:"access_token"`
	User        UserResponseLoginRequest `json:"user"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterUserRequestParams
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	args := database.CreateUserTXParams{
		CreateUserParams: database.CreateUserParams{
			Username:        req.Username,
			HashedPassword:  hashPassword,
			Email:           req.Email,
			FullName:        req.FullName,
			RoleID:          10,
			CodeVerifyEmail: pgtype.Text{String: utils.RandomCode(), Valid: true},
		},
		AfterCreate: func(u database.User) error {
			taskPayload := &mailer.UserReceive{
				Username:     u.Username,
				EmailAddress: u.Email,
				Code:         u.CodeVerifyEmail.String,
				Fullname:     u.FullName,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return server.taskDistributor.DistributeTaskSendVerifyAccount(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTX(ctx, args)
	if err != nil {
		errCodePgx := utils.ErrorCodePgxConstraint(err)
		if errCodePgx == utils.UniqueViolation {
			ctx.JSON(http.StatusConflict, gin.H{
				"error_unique": err.Error(),
			})
			return
		}

		if errCodePgx == utils.ForeignKeyViolation {
			ctx.JSON(http.StatusConflict, gin.H{
				"error_foreignkey": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, txResult.User)
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequestParams

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Identifier)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Khong tim thay nguoi dung nay")))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.ComparePassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Mat khau khong chinh xac")))
		return
	}

	if !user.IsVerifiedEmail {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("Tai khoan chua duoc kich hoat, vui long kiem tra: %s", user.Email)))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, int(user.RoleID), server.env.TIME_EXPIRED_TOKEN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loginResponseBody{
		AccessToken: accessToken,
		User: UserResponseLoginRequest{
			Username:          user.Username,
			Email:             user.Email,
			FullName:          user.FullName,
			IsVerifiedEmail:   user.IsVerifiedEmail,
			PasswordChangedAt: user.PasswordChangedAt.Time,
			CreatedAt:         user.CreatedAt.Time,
		},
	})
}
