-- UP
-- Manually Upsert Allison to guarantee she exists and is featured
INSERT INTO pets (
    id, 
    name, 
    date_of_birth, 
    breed, 
    sex, 
    status, 
    descriptions, 
    profile_settings,
    photos
) 
VALUES (
    2, 
    'Allison', 
    '2023-12-08', 
    'Domestic Shorthair', 
    'female', 
    'available',
    '{"spotlight": "Sweet Torti. A total lovebug who thrives as the only pet."}', 
    '{"isSpotlightFeatured": "true"}',
    '[]'
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    status = EXCLUDED.status,
    descriptions = pets.descriptions || EXCLUDED.descriptions,
    profile_settings = jsonb_set(COALESCE(pets.profile_settings, '{}'), '{isSpotlightFeatured}', '"true"');

-- Also ensure the others are still set (redundancy is fine)
UPDATE pets
SET profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"')
WHERE name IN ('Sally', 'Merry', 'Colby');

-- DOWN
