package main

import (
	"context"
	"log"

	"github.com/go-http-server/core/api"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	env, err := utils.LoadEnviromentVariables("./")
	if err != nil {
		log.Fatal("Cannot load enviroment variables: ", err)
	}

	pool, err := pgxpool.New(context.Background(), env.DB_SOURCE)
	if err != nil {
		log.Fatal("Cannot create pool to database: ", err)
	}
	store := database.NewStore(pool)

	server, err := api.NewServer(store, env)
	if err != nil {
		log.Fatal("Cannot create new server: ", err)
	}
	server.StartServer(env.HTTP_SERVER_ADDRESS)
}
