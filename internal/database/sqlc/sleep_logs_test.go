package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
	"time"
)

var user = NewCreateUserParams()

func TestCreateSleepLog(t *testing.T) {
	_, _ = testQueries.CreateUser(context.Background(), user)

	id, _ := uuid.NewUUID()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	sleepLog, err := testQueries.CreateSleepLog(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, arg.ID, sleepLog.ID)
	require.Equal(t, arg.StartTime, sleepLog.StartTime)
	require.Equal(t, arg.EndTime, sleepLog.EndTime)
	require.Equal(t, arg.Quality, sleepLog.Quality)
	require.NotZero(t, sleepLog.CreatedAt)
}

func TestGetSleepLogsByUserID(t *testing.T) {
	_, _ = testQueries.CreateUser(context.Background(), user)

	id, _ := uuid.NewUUID()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	createdSleepLog, err := testQueries.CreateSleepLog(context.Background(), arg)

	query := GetSleepLogsByUserIDParams{
		UserID: user.ID,
		Limit:  50,
		Offset: 0,
	}
	sleepLogs, err := testQueries.GetSleepLogsByUserID(context.Background(), query)

	idx := slices.IndexFunc(sleepLogs, func(log SleepLog) bool {
		return log.ID == createdSleepLog.ID
	})

	require.NoError(t, err)
	require.Equal(t, createdSleepLog.ID, sleepLogs[idx].ID)
	require.Equal(t, createdSleepLog.StartTime, sleepLogs[idx].StartTime)
	require.Equal(t, createdSleepLog.EndTime, sleepLogs[idx].EndTime)
	require.Equal(t, createdSleepLog.Quality, sleepLogs[idx].Quality)
	require.Equal(t, createdSleepLog.CreatedAt, sleepLogs[idx].CreatedAt)
}
