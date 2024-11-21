package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/utils"
)

type Server struct {
	store  database.Store
	router *gin.Engine
	env    utils.EnviromentVariables
}

func NewServer(store database.Store, env utils.EnviromentVariables) (*Server, error) {
	server := &Server{
		store: store,
		env:   env,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	if server.env.ENVIRONMENT == "product" {
		gin.DisableConsoleColor()
	}
	router := gin.Default()

	api := router.Group("/api")
	{
		noAuthRoute := api.Group("/")
		noAuthRoute.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
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
