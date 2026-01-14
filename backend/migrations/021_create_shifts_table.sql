CREATE TABLE IF NOT EXISTS shifts (
    id bigserial PRIMARY KEY,
    volunteer_id bigint NOT NULL REFERENCES volunteers(id) ON DELETE CASCADE,
    date date NOT NULL,
    start_time text NOT NULL,
    end_time text NOT NULL,
    role text NOT NULL,
    status text NOT NULL DEFAULT 'scheduled',
    notes text NOT NULL DEFAULT '',
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_shifts_volunteer_id ON shifts(volunteer_id);
CREATE INDEX IF NOT EXISTS idx_shifts_date ON shifts(date);

GRANT ALL PRIVILEGES ON TABLE shifts TO PUBLIC;
GRANT ALL PRIVILEGES ON SEQUENCE shifts_id_seq TO PUBLIC;
