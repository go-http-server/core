package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomRole(t *testing.T) {
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
