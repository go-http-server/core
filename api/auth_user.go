package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/utils"
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

	arg := database.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashPassword,
		Email:          req.Email,
		FullName:       req.FullName,
		RoleID:         10,
	}

	user, err := server.store.CreateUser(ctx, arg)
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
