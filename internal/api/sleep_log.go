package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"net/http"
	"time"
)

type createSleepLogRequest struct {
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	Quality   string    `json:"quality" binding:"required,oneof='Very Poor' 'Poor' 'Fair' 'Good' 'Very Good' 'Excellent'"`
}

func (s *Server) createSleepLog(ctx *gin.Context) {
	var req createSleepLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateSleepLogParams{
		ID:        uuid.New(),
		UserID:    authPayload.UserID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Quality:   req.Quality,
	}

	sleepLog, err := s.store.CreateSleepLog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sleepLog)
}

type getSleepLogs struct {
	pageNumber int32 `json:"pageNumber" binding:"min=1"`
	pageSize   int32 `json:"pageSize" binding:"min=1,max=100"`
}

func (logs *getSleepLogs) checkDefaults() {
	if logs.pageNumber == 0 {
		logs.pageNumber = 1
	}
	if logs.pageSize == 0 {
		logs.pageSize = 10
	}
}

// GetSleepLogsByUserID returns all sleep logs for a user
func (s *Server) getSleepLogsByUserID(ctx *gin.Context) {
	var req getSleepLogs
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.checkDefaults()

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetSleepLogsByUserIDParams{
		UserID: authPayload.UserID,
		Limit:  req.pageSize,
		Offset: (req.pageNumber - 1) * req.pageSize,
	}

	sleepLogs, err := s.store.GetSleepLogsByUserID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sleepLogs)
}
