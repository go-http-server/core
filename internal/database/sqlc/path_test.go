package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPath(t *testing.T) {
	path := utils.RandomString(6)
	pathName := utils.RandomString(8)
	pathDescription := utils.RandomString(12)

	pathFromDb, err := testStore.CreatePath(context.Background(), CreatePathParams{
		Path:            path,
		PathDescription: pgtype.Text{String: pathDescription, Valid: true},
		PathName:        pathName,
	})
	require.NoError(t, err)
}

func TestCreatePath(t *testing.T) {
	createRandomPath(t)
}
