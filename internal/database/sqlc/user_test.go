package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	arg := NewCreateUserParams()

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.CreatedAt)
}

func TestGetUserByUsername(t *testing.T) {
	t.Parallel()

	arg := NewCreateUserParams()

	_, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	user, err := testStore.GetUserByUsername(context.Background(), arg.Username)

	require.NoError(t, err)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.CreatedAt)
}
