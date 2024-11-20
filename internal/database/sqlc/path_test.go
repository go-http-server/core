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
	require.NotEmpty(t, pathFromDb)

	require.Equal(t, path, pathFromDb.Path)
	require.Equal(t, pathName, pathFromDb.PathName)
	require.Equal(t, pathDescription, pathFromDb.PathDescription.String)
}

func TestCreatePath(t *testing.T) {
	createRandomPath(t)
}

func TestListPaths(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomPath(t)
	}

	arg := ListPathsParams{
		Offset: 0,
		Limit:  5,
	}

	listPaths, err := testStore.ListPaths(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listPaths, 5)
}
