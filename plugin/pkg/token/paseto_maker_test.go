package token

import (
	"testing"
	"time"

	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePasetoAndVerifyTokenSuccess(t *testing.T) {
	maker := NewPasetoMaker()

	username := utils.RandomString(6)
	roleId := utils.RandomInt(1, 10)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, roleId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, payload)

	require.Equal(t, username, payload.Username)
	require.Equal(t, roleId, payload.RoleId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
	require.WithinDuration(t, payload.IssuedAt, payload.ExpiredAt, 2*duration)
}

func TestExpiredToken(t *testing.T) {
	maker := NewPasetoMaker()

	token, err := maker.CreateToken(utils.RandomString(6), utils.RandomInt(1, 10), -time.Minute)
	require.NoError(t, err)
	require.NotNil(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, utils.TOKEN_EXPIRED)
	require.Nil(t, payload)
	require.Empty(t, payload)
}
