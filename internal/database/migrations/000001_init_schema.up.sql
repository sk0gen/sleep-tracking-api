-- Create the users table
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON users (username);

CREATE TABLE sleep_logs (
                            id UUID PRIMARY KEY,
                            user_id UUID,
                            start_time TIMESTAMP NOT NULL,
                            end_time TIMESTAMP NOT NULL,
                            quality VARCHAR(50) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE sleep_logs ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE INDEX ON sleep_logs (user_id);