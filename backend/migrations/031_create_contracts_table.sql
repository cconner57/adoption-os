-- Up Migration
CREATE TABLE IF NOT EXISTS contracts (
    token text PRIMARY KEY,
    application_id bigint NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    type text NOT NULL, -- 'STANDARD', 'FOSTER_TO_ADOPT'
    expires_at timestamp(0) with time zone NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_contracts_application_id ON contracts(application_id);
