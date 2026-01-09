-- UP
-- Try to find Allison more robustly (case-insensitive, wildcards for whitespace)
UPDATE pets
SET profile_settings = jsonb_set(
    COALESCE(profile_settings, '{}'), 
    '{isSpotlightFeatured}', 
    '"true"'
)
WHERE name ILIKE '%Allison%';

-- DOWN
