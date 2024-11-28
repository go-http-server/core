package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store database.Store) *Server {
	env := utils.EnviromentVariables{
		DB_SOURCE:           "",
		HTTP_SERVER_ADDRESS: "",
		TIME_EXPIRED_TOKEN:  30 * time.Minute,
	}

	testServer, err := NewServer(store, env, nil)
	require.NoError(t, err)
	require.NotEmpty(t, testServer)

	return testServer
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
