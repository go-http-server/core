package api

import (
	"context"
	"errors"
	"net/http"

	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/token"
	"github.com/go-http-server/core/utils"
	"github.com/go-http-server/core/worker"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	store           database.Store
	router          *gin.Engine
	env             utils.EnvironmentVariables
	tokenMaker      token.TokenMaker
	taskDistributor worker.TaskDistributor
}

func NewServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	store database.Store,
	env utils.EnvironmentVariables,
	taskDistributor worker.TaskDistributor,
) (*Server, error) {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	parser := paseto.NewParserWithoutExpiryCheck()
	tokenMaker := token.NewPasetoMaker(privateKey, parser)

	server := &Server{
		store:           store,
		env:             env,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	if server.env.ENVIRONMENT == utils.ProductionEnvironment {
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

func (server *Server) StartServer(ctx context.Context, waitGroup *errgroup.Group, address string) {
	httpServer := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("Start HTTP Server at %s", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("HTTP Server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Graceful shutdown HTTP server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Failed to shutdown HTTP server")
			return err
		}

		log.Info().Msg("HTTP Server was stopped")
		return nil
	})
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}
