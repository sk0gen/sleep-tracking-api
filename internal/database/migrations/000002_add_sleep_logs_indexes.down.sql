ALTER TABLE sleep_logs
    DROP CONSTRAINT sleep_logs_overlaps;

DROP INDEX idx_sleep_logs_start_time;