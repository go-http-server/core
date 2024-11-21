package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPath(t *testing.T) Path {
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

	return pathFromDb
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

func TestUpdatePath(t *testing.T) {
	randomPath := createRandomPath(t)
	pathName := utils.RandomString(6)
	path := utils.RandomString(6)
	pathDescription := utils.RandomString(12)

	arg := UpdatePathParams{
		ID:              randomPath.ID,
		PathName:        pgtype.Text{String: pathName, Valid: true},
		Path:            pgtype.Text{String: path, Valid: true},
		PathDescription: pgtype.Text{String: pathDescription, Valid: true},
	}

	pathStore, err := testStore.UpdatePath(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pathStore)
	require.Equal(t, pathName, pathStore.PathName)
	require.Equal(t, path, pathStore.Path)
	require.Equal(t, pathDescription, pathStore.PathDescription.String)
}
