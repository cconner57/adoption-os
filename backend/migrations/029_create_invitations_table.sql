CREATE TABLE IF NOT EXISTS invitations (
    token text PRIMARY KEY,
    email text NOT NULL,
    role text NOT NULL,
    expires_at TIMESTAMP(0) WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
