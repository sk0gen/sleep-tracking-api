package token

import (
	"github.com/sk0gen/sleep-tracking-api/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJWTMaker(t *testing.T) {
	maker := NewJWTMaker(util.RandomString(32))

	username := util.RandomString(6)
	duration := time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.Id)
	require.Equal(t, username, payload.Username)
}

func TestExpiredJWTToken(t *testing.T) {
	maker := NewJWTMaker(util.RandomString(32))

	username := util.RandomString(6)
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
