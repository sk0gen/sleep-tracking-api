package db

import (
	"github.com/google/uuid"
	"github.com/sk0gen/sleep-tracking-api/util"
)

func NewCreateUserParams() CreateUserParams {
	id := uuid.New()

	password := util.RandomString(6)
	passwordHash, _ := util.HashPassword(password)

	return CreateUserParams{
		ID:           id,
		Username:     uuid.NewString(),
		PasswordHash: passwordHash,
	}
}
