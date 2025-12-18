-- UP
-- We use IF NOT EXISTS so this script won't crash if you've already run some of these commands manually.

ALTER TABLE pets 
    ADD COLUMN IF NOT EXISTS slug TEXT,
    ADD COLUMN IF NOT EXISTS litter_name TEXT,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW();

ALTER TABLE pets 
    ADD COLUMN IF NOT EXISTS physical JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS behavior JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS medical JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS descriptions JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS details JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS adoption JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS foster JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS returned JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS sponsored JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS photos JSONB DEFAULT '[]',
    ADD COLUMN IF NOT EXISTS profile_settings JSONB DEFAULT '{}';

-- DOWN
-- This removes the columns if we ever need to roll back.

ALTER TABLE pets 
    DROP COLUMN IF EXISTS slug,
    DROP COLUMN IF EXISTS litter_name,
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at,
    DROP COLUMN IF EXISTS physical,
    DROP COLUMN IF EXISTS behavior,
    DROP COLUMN IF EXISTS medical,
    DROP COLUMN IF EXISTS descriptions,
    DROP COLUMN IF EXISTS details,
    DROP COLUMN IF EXISTS adoption,
    DROP COLUMN IF EXISTS foster,
    DROP COLUMN IF EXISTS returned,
    DROP COLUMN IF EXISTS sponsored,
    DROP COLUMN IF EXISTS photos,
    DROP COLUMN IF EXISTS profile_settings;