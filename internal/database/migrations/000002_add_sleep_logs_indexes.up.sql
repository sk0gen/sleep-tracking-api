ALTER TABLE sleep_logs
    ADD CONSTRAINT sleep_logs_overlaps
    EXCLUDE USING gist (tsrange(start_time, end_time) WITH &&);

CREATE INDEX idx_sleep_logs_start_time ON sleep_logs (start_time);