package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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

	sleepLog, err := testQueries.CreateSleepLog(context.Background(), arg)

	sleepLogs, err := testQueries.GetSleepLogsByUserID(context.Background(), user.ID)

	require.NoError(t, err)
	require.Equal(t, 1, len(sleepLogs))
	require.Equal(t, sleepLog.ID, sleepLogs[0].ID)
	require.Equal(t, sleepLog.StartTime, sleepLogs[0].StartTime)
	require.Equal(t, sleepLog.EndTime, sleepLogs[0].EndTime)
	require.Equal(t, sleepLog.Quality, sleepLogs[0].Quality)
	require.Equal(t, sleepLog.CreatedAt, sleepLogs[0].CreatedAt)
}
