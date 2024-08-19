package db

import "github.com/google/uuid"

func NewCreateUserParams() CreateUserParams {
	id, _ := uuid.NewUUID()

	return CreateUserParams{
		ID:           id,
		Username:     uuid.NewString(),
		PasswordHash: "passwordHash",
	}
}
