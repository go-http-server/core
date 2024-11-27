package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/go-http-server/core/utils"
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
			subject := "[Go core] Kích hoạt tài khoản"
			pathTemplate := "./templates/verify_email.html"
			to := mailer.UserReceive{
				Username:     u.Username,
				EmailAddress: u.Email,
				Code:         u.CodeVerifyEmail.String,
				Fullname:     u.FullName,
			}
			return server.emailSender.SendWithTemplate(subject, pathTemplate, to, []string{})
		},
	}

	user, err := server.store.CreateUserTX(ctx, args)
	if err != nil {
		errCodePgx := utils.ErrorCodePgxConstraint(err)
		if errCodePgx == utils.UniqueViolation {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error_unique": err.Error(),
			})
			return
		}

		if errCodePgx == utils.ForeignKeyViolation {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error_foreignkey": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (server *Server) LoginUser(ctx *gin.Context) {
}
