package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
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

	arg := db.CreateSleepLogParams{
		ID:        uuid.New(),
		UserID:    uuid.New(), //TODO: Get from claims
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Quality:   req.Quality,
	}

	sleepLog, err := s.database.CreateSleepLog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sleepLog)
}

type getSleepLogs struct {
	pageNumber int32 `json:"pageNumber" binding:"required,min=1"`
	pageSize   int32 `json:"pageSize" binding:"required,min=1,max=100"`
}

// GetSleepLogsByUserID returns all sleep logs for a user
func (s *Server) getSleepLogsByUserID(ctx *gin.Context) {
	var req getSleepLogs
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(ctx.Param("userID")) // TODO: GET userid from claims
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetSleepLogsByUserIDParams{
		UserID: userID,
		Limit:  req.pageSize,
		Offset: (req.pageNumber - 1) * req.pageSize,
	}

	sleepLogs, err := s.database.GetSleepLogsByUserID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sleepLogs)
}
