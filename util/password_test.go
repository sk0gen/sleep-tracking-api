package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword)
	require.Error(t, err)

	hashedPassword2, err := HashPassword(password)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}
