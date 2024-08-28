package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
	"time"
)

func TestCreateSleepLog(t *testing.T) {
	t.Parallel()
	user, _ := testStore.CreateUser(context.Background(), NewCreateUserParams())

	id := uuid.New()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	sleepLog, err := testStore.CreateSleepLog(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, arg.ID, sleepLog.ID)
	require.Equal(t, arg.StartTime, sleepLog.StartTime)
	require.Equal(t, arg.EndTime, sleepLog.EndTime)
	require.Equal(t, arg.Quality, sleepLog.Quality)
	require.NotZero(t, sleepLog.CreatedAt)
}

func TestGetSleepLogsByUserID(t *testing.T) {
	t.Parallel()

	user, _ := testStore.CreateUser(context.Background(), NewCreateUserParams())

	id := uuid.New()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 1, 2, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	createdSleepLog, err := testStore.CreateSleepLog(context.Background(), arg)

	query := GetSleepLogsByUserIDParams{
		UserID: user.ID,
		Limit:  50,
		Offset: 0,
	}
	sleepLogs, err := testStore.GetSleepLogsByUserID(context.Background(), query)

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

func TestDeleteSleepLogById(t *testing.T) {
	t.Parallel()

	user, _ := testStore.CreateUser(context.Background(), NewCreateUserParams())

	id := uuid.New()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 1, 3, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	_, err := testStore.CreateSleepLog(context.Background(), arg)
	require.NoError(t, err)

	query := GetSleepLogsByUserIDParams{
		UserID: user.ID,
		Limit:  50,
		Offset: 0,
	}
	sleepLogs, err := testStore.GetSleepLogsByUserID(context.Background(), query)
	require.NoError(t, err)
	require.Len(t, sleepLogs, 1)

	deleteQuery := DeleteSleepLogByIDParams{
		ID:     id,
		UserID: user.ID,
	}

	err = testStore.DeleteSleepLogByID(context.Background(), deleteQuery)
	require.NoError(t, err)

	sleepLogs, err = testStore.GetSleepLogsByUserID(context.Background(), query)
	require.NoError(t, err)
	require.Len(t, sleepLogs, 0)
}

func TestUpdateSleepLogById(t *testing.T) {
	t.Parallel()

	user, _ := testStore.CreateUser(context.Background(), NewCreateUserParams())

	id := uuid.New()

	arg := CreateSleepLogParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 2, 3, 8, 0, 0, 0, time.UTC),
		Quality:   "Good",
	}

	_, err := testStore.CreateSleepLog(context.Background(), arg)
	require.NoError(t, err)

	query := GetSleepLogsByUserIDParams{
		UserID: user.ID,
		Limit:  50,
		Offset: 0,
	}
	sleepLogs, err := testStore.GetSleepLogsByUserID(context.Background(), query)
	require.NoError(t, err)
	require.Len(t, sleepLogs, 1)

	updateSleepLog := UpdateSleepLogByIdParams{
		ID:        id,
		UserID:    user.ID,
		StartTime: time.Date(2021, 2, 4, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2021, 2, 4, 8, 0, 0, 0, time.UTC),
		Quality:   "Bad",
	}

	err = testStore.UpdateSleepLogById(context.Background(), updateSleepLog)
	require.NoError(t, err)

	sleepLogs, err = testStore.GetSleepLogsByUserID(context.Background(), query)
	require.NoError(t, err)
	require.Len(t, sleepLogs, 1)
	require.Equal(t, updateSleepLog.StartTime, sleepLogs[0].StartTime)
	require.Equal(t, updateSleepLog.EndTime, sleepLogs[0].EndTime)
	require.Equal(t, updateSleepLog.Quality, sleepLogs[0].Quality)
}
