package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomRole(t *testing.T) Role {
	roleName := utils.RandomString(32)
	roleDescription := utils.RandomString(32)

	arg := CreateRoleParams{
		RoleName:    roleName,
		Description: pgtype.Text{String: roleDescription, Valid: true},
	}

	role, err := testStore.CreateRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, role)

	require.Equal(t, roleName, role.RoleName)
	require.Equal(t, roleDescription, role.Description.String)

	return role
}

func TestCreateRole(t *testing.T) {
	createRandomRole(t)
}

func TestListRoles(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomRole(t)
	}

	arg := GetRolesParams{
		Limit:  5,
		Offset: 0,
	}

	listRoles, err := testStore.GetRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listRoles, 5)
}

func TestUpdateRole(t *testing.T) {
	randomRole := createRandomRole(t)
	roleName := utils.RandomString(6)
	roleDescription := utils.RandomString(12)

	arg := UpdateRoleParams{
		ID:          randomRole.ID,
		RoleName:    pgtype.Text{String: roleName, Valid: true},
		Description: pgtype.Text{String: roleDescription, Valid: true},
	}

	roleUpdated, err := testStore.UpdateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roleUpdated)
	require.Equal(t, roleName, roleUpdated.RoleName)
	require.Equal(t, roleDescription, roleUpdated.Description.String)
}

func TestGetOneRole(t *testing.T) {
	randomRole := createRandomRole(t)

	storeRole, err := testStore.GetOneRole(context.Background(), randomRole.ID)
	require.NoError(t, err)
	require.NotEmpty(t, storeRole)

	require.Equal(t, randomRole.ID, storeRole.ID)
	require.Equal(t, randomRole.RoleName, storeRole.RoleName)
	require.Equal(t, randomRole.Description, storeRole.Description)
}
