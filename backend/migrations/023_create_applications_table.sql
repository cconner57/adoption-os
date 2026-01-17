CREATE TABLE IF NOT EXISTS applications (
    id bigserial PRIMARY KEY,
    type text NOT NULL, -- 'volunteer', 'adoption', 'surrender'
    status text NOT NULL DEFAULT 'pending', -- 'pending', 'approved', 'denied', 'needs_info'
    data jsonb NOT NULL DEFAULT '{}'::jsonb,
    original_html text, -- Stores the generated HTML of the email for "View Original"
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);

-- Add indexes
CREATE INDEX IF NOT EXISTS idx_applications_type ON applications(type);
CREATE INDEX IF NOT EXISTS idx_applications_status ON applications(status);
