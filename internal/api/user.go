package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/util"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum" example:"patient1"`
	Password string `json:"password" binding:"required,min=10,strongpassword" example:"Str0ngP@ssw0rd!"`
}

type userResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

// @BasePath /api/v1
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param input body createUserRequest true "User data"
// @Success 200 {object} userResponse
// @Failure 400 {object} errResponse
// @Router /auth/register [post]
func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
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
		ID:           uuid.New(),
		Username:     req.Username,
		PasswordHash: hashedPassword,
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required" example:"patient1"`
	Password string `json:"password" binding:"required" example:"Str0ngP@ssw0rd!"`
}

type loginUserResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param input body loginUserRequest true "User data"
// @Success 200 {object} loginUserResponse
// @Failure 400 {object} errResponse
// @Router /auth/login [post]
func (s *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.PasswordHash)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := s.jwtMaker.CreateToken(
		user.ID,
		s.config.AuthConfig.JWTTokenExpiration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		Token: accessToken,
		User:  newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
