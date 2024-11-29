package database

import (
	"context"
	"testing"
	"time"

	"github.com/go-http-server/core/utils"
	"github.com/jackc/pgx/v5/pgtype"
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

func TestGetUser(t *testing.T) {
	randomUser := createRandomUser(t)

	userFromDB, err := testStore.GetUser(context.Background(), randomUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userFromDB)
	require.Equal(t, randomUser.Username, userFromDB.Username)
	require.Equal(t, randomUser.CodeVerifyEmail, userFromDB.CodeVerifyEmail)
	require.Equal(t, randomUser.CodeResetPassword, userFromDB.CodeResetPassword)
	require.Equal(t, randomUser.HashedPassword, userFromDB.HashedPassword)
	require.Equal(t, randomUser.Email, userFromDB.Email)

	require.Equal(t, randomUser.IsVerifiedEmail, userFromDB.IsVerifiedEmail)
	require.Equal(t, randomUser.FullName, userFromDB.FullName)
	require.Equal(t, randomUser.RoleID, userFromDB.RoleID)
	require.Equal(t, randomUser.Token, userFromDB.Token)
	require.Equal(t, randomUser.PasswordChangedAt, userFromDB.PasswordChangedAt)
}

func TestUpdateUser(t *testing.T) {
	randomUser := createRandomUser(t)

	updateFullName := utils.RandomString(18)
	updateEmail := utils.RandomEmail()

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		FullName: pgtype.Text{String: updateFullName, Valid: true},
		Email:    pgtype.Text{String: updateEmail, Valid: true},
		Username: randomUser.Username,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, updateFullName, updatedUser.FullName)
	require.Equal(t, updateEmail, updatedUser.Email)
	require.Equal(t, randomUser.Username, updatedUser.Username)

	require.Equal(t, randomUser.Token, updatedUser.Token)
	require.Equal(t, randomUser.HashedPassword, randomUser.HashedPassword)
	require.Equal(t, randomUser.RoleID, updatedUser.RoleID)
	require.WithinDuration(t, randomUser.CreatedAt.Time, updatedUser.CreatedAt.Time, time.Minute)
}
