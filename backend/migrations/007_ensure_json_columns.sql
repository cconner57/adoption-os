-- UP
ALTER TABLE pets 
    ADD COLUMN IF NOT EXISTS descriptions JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS photos JSONB DEFAULT '[]',
    ADD COLUMN IF NOT EXISTS profile_settings JSONB DEFAULT '{}';

-- DOWN
-- No-op or drop if strictly needed, but better to keep for safety
