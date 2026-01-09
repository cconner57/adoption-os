-- UP
-- 1. Reset all pets spotlight status to false
UPDATE pets
SET profile_settings = jsonb_set(
    COALESCE(profile_settings, '{}'), 
    '{isSpotlightFeatured}', 
    '"false"'
);

-- 2. Set spotlight TRUE for specific pets
-- Using ILIKE or LOWER to be Case-Insensitive safe, though names are likely capitalized.
UPDATE pets
SET profile_settings = jsonb_set(
    profile_settings, 
    '{isSpotlightFeatured}', 
    '"true"'
)
WHERE name IN ('Sally', 'Merry', 'Colby', 'Allison');

-- DOWN
-- No-op, or revert to random 4? safely ignore.
