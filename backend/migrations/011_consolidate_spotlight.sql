-- UP
-- Ensure all 4 target pets are spotlighted, correcting for any past misses
UPDATE pets
SET profile_settings = jsonb_set(
    COALESCE(profile_settings, '{}'), 
    '{isSpotlightFeatured}', 
    '"true"'
)
WHERE name ILIKE '%Allison%' 
   OR name ILIKE '%Sally%' 
   OR name ILIKE '%Merry%' 
   OR name ILIKE '%Colby%';

-- DOWN
