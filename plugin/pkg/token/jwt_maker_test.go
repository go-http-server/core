package token

import (
	"testing"
	"time"

	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTToken(t *testing.T) {
	tokenMaker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)

	username := utils.RandomString(6)
	roleId := utils.RandomInt(1, 10)

	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := tokenMaker.CreateToken(username, roleId, duration)
	require.NoError(t, err)
	require.NotNil(t, token)

	payload, err := tokenMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, username, payload.Username)
	require.Equal(t, roleId, payload.RoleId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Minute)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Minute)
}

func TestExpiredToken(t *testing.T) {
	tokenMaker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)

	username := utils.RandomString(6)
	roleId := utils.RandomInt(1, 10)

	duration := -time.Minute

	token, err := tokenMaker.CreateToken(username, roleId, duration)
	require.NoError(t, err)
	require.NotNil(t, token)

	payload, err := tokenMaker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, utils.TOKEN_EXPIRED)
	require.Nil(t, payload)
}
