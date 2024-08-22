package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/util"
	"net/http"
	"time"
)

type CreateUserParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=10"`
}

type UserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: hashedPassword,
	}

	user, err := s.database.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}
