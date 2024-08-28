package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/pagination"
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

	if req.EndTime.Before(req.StartTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Sleep end time must be after start time"})
		return
	}

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

// GetSleepLogsByUserID returns all sleep logs for a user
func (s *Server) getSleepLogs(ctx *gin.Context) {
	var req pagination.PaginatedRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.CheckDefaults()

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetSleepLogsByUserIDParams{
		UserID: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}

	sleepLogs, err := s.store.GetSleepLogsByUserID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	sleepLogsCount, err := s.store.GetSleepLogCountByUserID(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	paginatedResponse := pagination.PaginatedResponse[db.SleepLog]{
		Results:    sleepLogs,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		TotalItems: sleepLogsCount,
	}

	ctx.JSON(http.StatusOK, paginatedResponse)
}

// DeleteSleepLogByID deletes a sleep log by ID
func (s *Server) deleteSleepLogByID(ctx *gin.Context) {
	var idRequest idUriRequest
	if err := ctx.ShouldBindUri(&idRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	ID, err := uuid.Parse(idRequest.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteSleepLogByIDParams{
		ID:     ID,
		UserID: authPayload.UserID,
	}

	err = s.store.DeleteSleepLogByID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type updateSleepLogRequest struct {
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	Quality   string    `json:"quality" binding:"required,oneof='Very Poor' 'Poor' 'Fair' 'Good' 'Very Good' 'Excellent'"`
}

func (s *Server) updateSleepLogByID(ctx *gin.Context) {
	var idRequest idUriRequest
	if err := ctx.ShouldBindUri(&idRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateSleepLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	ID, err := uuid.Parse(idRequest.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.EndTime.Before(req.StartTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Sleep end time must be after start time"})
		return
	}

	arg := db.UpdateSleepLogByIdParams{
		ID:        ID,
		UserID:    authPayload.UserID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Quality:   req.Quality,
	}

	err = s.store.UpdateSleepLogById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
