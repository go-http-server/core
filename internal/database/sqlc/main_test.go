package database

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	env, err := utils.LoadEnviromentVariables("../../")
	if err != nil {
		log.Fatal("Cannot load enviroment variables file: ", err)
	}

	connectionPool, err := pgxpool.New(context.Background(), env.DB_SOURCE)
	if err != nil {
		log.Fatal("Cannot create connection pool into database: ", err)
	}
	defer connectionPool.Close()
	testStore = NewStore(connectionPool)
	os.Exit(m.Run())
}
