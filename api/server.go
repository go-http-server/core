package api

import (
	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/go-http-server/core/plugin/pkg/token"
	"github.com/go-http-server/core/utils"
	"github.com/go-http-server/core/worker"
)

type Server struct {
	store           database.Store
	router          *gin.Engine
	env             utils.EnviromentVariables
	tokenMaker      token.TokenMaker
	emailSender     mailer.EmailSender
	taskDistributor worker.TaskDistributor
}

func NewServer(store database.Store, env utils.EnviromentVariables, taskDistributor worker.TaskDistributor) (*Server, error) {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	parser := paseto.NewParserWithoutExpiryCheck()
	tokenMaker := token.NewPasetoMaker(privateKey, parser)

	emailSender := mailer.NewGmailSender(env.EMAIL_USERNAME_SENDER, env.EMAIL_ADDRESS_SENDER, env.EMAIL_PASSWORD_SENDER)

	server := &Server{
		store:           store,
		env:             env,
		tokenMaker:      tokenMaker,
		emailSender:     emailSender,
		taskDistributor: taskDistributor,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	if server.env.ENVIRONMENT == "product" {
		gin.DisableConsoleColor()
	}
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		noRequiredAuthRoute := api.Group("/")
		noRequiredAuthRoute.POST("/auth/register", server.RegisterUser)
		noRequiredAuthRoute.POST("/auth/login", server.LoginUser)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"api":     "Core api",
			"author":  "phamnam2003",
			"version": "0.0.1",
			"email":   "namphamhai7@gmail.com",
		})
	})

	server.router = router
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}
