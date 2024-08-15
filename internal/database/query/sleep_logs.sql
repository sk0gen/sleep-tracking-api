-- name: CreateSleepLog :one
INSERT INTO sleep_logs (id, user_id, start_time, end_time, quality, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;