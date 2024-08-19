package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := NewCreateUserParams()

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.CreatedAt)
}

func TestGetUserByUsername(t *testing.T) {
	arg := NewCreateUserParams()

	_, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	user, err := testQueries.GetUserByUsername(context.Background(), arg.Username)

	require.NoError(t, err)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.CreatedAt)
}
