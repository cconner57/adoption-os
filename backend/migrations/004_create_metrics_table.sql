CREATE TABLE IF NOT EXISTS metrics (
    id bigserial PRIMARY KEY,
    event_type text NOT NULL,
    event_data jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_metrics_event_type ON metrics(event_type);
CREATE INDEX IF NOT EXISTS idx_metrics_created_at ON metrics(created_at);
