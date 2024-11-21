package database

import (
	"context"
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	username := utils.RandomString(6)
	hashedPassword := utils.RandomString(18)
	email := utils.RandomEmail()
	fullName := utils.RandomString(24)

	role, err := testStore.CreateRole(context.Background(), CreateRoleParams{
		RoleName: utils.RandomString(8),
	})
	require.NoError(t, err)

	user, err := testStore.CreateUser(context.Background(), CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		FullName:       fullName,
		RoleID:         role.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, username, user.Username)
	require.Equal(t, hashedPassword, user.HashedPassword)
	require.Equal(t, email, user.Email)
	require.Equal(t, fullName, user.FullName)

	require.Empty(t, user.CodeResetPassword)
	require.Empty(t, user.CodeVerifyEmail)
	require.Empty(t, user.Token)
	require.True(t, user.IsVerifiedEmail == false)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
