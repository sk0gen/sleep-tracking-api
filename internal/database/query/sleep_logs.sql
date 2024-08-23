-- name: CreateSleepLog :one
INSERT INTO sleep_logs (id, user_id, start_time, end_time, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetSleepLogsByUserID :many
SELECT *
FROM sleep_logs
WHERE user_id = $1
order by start_time desc
LIMIT $2
OFFSET $3;

-- name: DeleteSleepLogByID :exec
DELETE FROM sleep_logs
WHERE id = $1 AND user_id = $2;