-- We can drop and recreate since this is a new feature
DROP TABLE IF EXISTS sessions;

CREATE TABLE sessions (
    token text PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    ip inet NOT NULL,
    user_agent text NOT NULL,
    expiry timestamptz NOT NULL
);

CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
