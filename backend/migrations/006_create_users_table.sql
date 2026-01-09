CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password_hash text NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);

-- Seed Admin User
-- Password: temp_password_123
INSERT INTO users (id, name, email, password_hash, activated)
VALUES (
    '00000000-0000-0000-0000-000000000001',
    'Admin',
    'admin@idohr.org',
    '$argon2id$v=19$m=65536,t=1,p=4$Rcdetsf4J/E1DJopRJ26NQ$cjMloG5y/MMuytBsk+JGpBr4pq3cHZwjPs5iCtqp9As',
    true
)
ON CONFLICT (email) DO NOTHING;
