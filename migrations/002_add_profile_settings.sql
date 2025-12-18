-- UP
ALTER TABLE pets ADD COLUMN profile_settings JSONB DEFAULT '{}';

-- DOWN (In case we need to undo it later)
-- ALTER TABLE pets DROP COLUMN profile_settings;