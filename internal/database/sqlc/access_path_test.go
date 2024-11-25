package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAccessPath(t *testing.T) {
	role, err := testStore.CreateRole(context.Background(), CreateRoleParams{
		RoleName:    utils.RandomString(12),
		Description: pgtype.Text{String: utils.RandomString(24)},
	})
	require.NoError(t, err)

	path, err := testStore.CreatePath(context.Background(), CreatePathParams{
		PathName: utils.RandomString(12),
		Path:     utils.RandomString(6),
	})
	require.NoError(t, err)

	accessPath, err := testStore.CreateAccessPath(context.Background(), CreateAccessPathParams{
		RoleID: role.ID,
		PathID: path.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, accessPath)
	require.Equal(t, role.ID, accessPath.RoleID)
	require.Equal(t, path.ID, accessPath.PathID)
}

func TestCreateAccessPath(t *testing.T) {
	createRandomAccessPath(t)
}
