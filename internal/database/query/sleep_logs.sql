-- name: CreateSleepLog :one
INSERT INTO sleep_logs (id, user_id, start_time, end_time, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetSleepLogsByUserID :many
SELECT *
FROM sleep_logs
WHERE user_id = $1;