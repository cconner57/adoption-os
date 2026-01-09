-- UP
-- We missed these in the previous force-fix. Adding them now to stop "column does not exist" errors.

ALTER TABLE pets 
    ADD COLUMN IF NOT EXISTS slug TEXT,
    ADD COLUMN IF NOT EXISTS litter_name TEXT,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS physical JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS behavior JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS medical JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS details JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS adoption JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS foster JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS returned JSONB DEFAULT '{}',
    ADD COLUMN IF NOT EXISTS sponsored JSONB DEFAULT '{}';

-- DOWN
