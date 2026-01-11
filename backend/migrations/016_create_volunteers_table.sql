CREATE TABLE IF NOT EXISTS volunteers (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    address text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    zip text NOT NULL DEFAULT '',
    role text NOT NULL DEFAULT 'Tier 1',
    status text NOT NULL DEFAULT 'active',
    bio text NOT NULL DEFAULT '',
    photo_url text NOT NULL DEFAULT '',
    reliability_score integer NOT NULL DEFAULT 100,
    total_hours integer NOT NULL DEFAULT 0,
    streak integer NOT NULL DEFAULT 0,
    join_date date NOT NULL DEFAULT CURRENT_DATE,
    allergies boolean NOT NULL DEFAULT false,
    skills text[] NOT NULL DEFAULT '{}',
    position_preferences text[] NOT NULL DEFAULT '{}',
    availability text[] NOT NULL DEFAULT '{}',
    badges text[] NOT NULL DEFAULT '{}',
    version integer NOT NULL DEFAULT 1
);

CREATE INDEX IF NOT EXISTS volunteers_email_idx ON volunteers(email);
CREATE INDEX IF NOT EXISTS volunteers_status_idx ON volunteers(status);
