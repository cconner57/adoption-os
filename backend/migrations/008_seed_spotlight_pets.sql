-- UP
-- Select 4 pets and enable spotlight for them
WITH target_pets AS (
    SELECT id FROM pets LIMIT 4
)
UPDATE pets
SET profile_settings = jsonb_set(
    COALESCE(profile_settings, '{}'), 
    '{isSpotlightFeatured}', 
    '"true"'
)
WHERE id IN (SELECT id FROM target_pets);

-- DOWN
-- Optional: Disable spotlight (reset)
-- UPDATE pets SET profile_settings = profile_settings - 'isSpotlightFeatured';
