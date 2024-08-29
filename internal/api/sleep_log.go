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
	StartTime time.Time `json:"startTime" binding:"required" example:"2020-01-01T22:00:00Z"`
	EndTime   time.Time `json:"endTime" binding:"required" example:"2020-01-02T08:00:00Z"`
	Quality   string    `json:"quality" binding:"required,oneof='Very Poor' 'Poor' 'Fair' 'Good' 'Very Good' 'Excellent'"`
}

type sleepLogResponse struct {
	ID        uuid.UUID `json:"id"`
	StartTime time.Time `json:"startTime" binding:"required" example:"2020-01-01T22:00:00Z"`
	EndTime   time.Time `json:"endTime" binding:"required" example:"2020-01-02T08:00:00Z"`
	Quality   string    `json:"quality" example:"Good"`
	CreatedAt time.Time `json:"createdAt" example:"2024-01-01T22:22:22Z"`
}

func newSleepLogResponse(sleepLog db.SleepLog) sleepLogResponse {
	return sleepLogResponse{
		ID:        sleepLog.ID,
		StartTime: sleepLog.StartTime,
		EndTime:   sleepLog.EndTime,
		Quality:   sleepLog.Quality,
		CreatedAt: sleepLog.CreatedAt,
	}
}

// @Summary Create sleep log
// @Description Create a new sleep log
// @Tags sleep-logs
// @Accept json
// @Produce json
// @Param input body createSleepLogRequest true "Sleep log data"
// @Success 200 {object} sleepLogResponse
// @Failure 400 {object} errResponse
// @Router /sleep-logs [post]
// @Security Bearer
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

	ctx.JSON(http.StatusOK, newSleepLogResponse(sleepLog))
}

// @Summary Get sleep logs
// @Description Get sleep logs
// @Tags sleep-logs
// @Accept json
// @Produce json
// @Param input query pagination.PaginatedRequest false "Pagination"
// @Success 200 {object} pagination.PaginatedResponse[sleepLogResponse]
// @Failure 400 {object} errResponse
// @Router /sleep-logs [get]
// @Security Bearer
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

	sleepLogsResponse := make([]sleepLogResponse, len(sleepLogs))
	for i, sleepLog := range sleepLogs {
		sleepLogsResponse[i] = newSleepLogResponse(sleepLog)
	}

	paginatedResponse := pagination.PaginatedResponse[sleepLogResponse]{
		Results:    sleepLogsResponse,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		TotalItems: sleepLogsCount,
	}

	ctx.JSON(http.StatusOK, paginatedResponse)
}

// @Summary Deletes sleep log
// @Description Deletes sleep logs
// @Tags sleep-logs
// @Accept json
// @Produce json
// @Param id path string true "Sleep log ID"
// @Success 204
// @Failure 400 {object} errResponse
// @Router /sleep-logs/{id} [delete]
// @Security Bearer
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
	StartTime time.Time `json:"startTime" binding:"required" example:"2020-01-01T22:00:00Z"`
	EndTime   time.Time `json:"endTime" binding:"required" example:"2020-01-02T08:00:00Z"`
	Quality   string    `json:"quality" binding:"required,oneof='Very Poor' 'Poor' 'Fair' 'Good' 'Very Good' 'Excellent'"`
}

// @Summary Updates sleep log
// @Description Updates sleep logs
// @Tags sleep-logs
// @Accept json
// @Produce json
// @Param id path string true "Sleep log ID"
// @Param input body updateSleepLogRequest true "Sleep log data"
// @Success 204
// @Failure 400 {object} errResponse
// @Router /sleep-logs/{id} [put]
// @Security Bearer
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
