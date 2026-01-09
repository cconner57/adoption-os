-- UP
-- Manually Upsert Allison to guarantee she exists and is featured (UUID compatible)
DO $$
DECLARE
    target_id uuid;
BEGIN
    -- 1. Try to find Allison by name
    SELECT id INTO target_id FROM pets WHERE name = 'Allison' LIMIT 1;

    IF target_id IS NOT NULL THEN
        -- 2. Update existing
        UPDATE pets SET 
            profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"'),
            descriptions = jsonb_set(COALESCE(descriptions, '{}'), '{spotlight}', '"Sweet Torti. A total lovebug who thrives as the only pet."'),
            status = 'available'
        WHERE id = target_id;
    ELSE
        -- 3. Insert new with generated UUID
        -- We use gen_random_uuid() which is built-in for PG 13+. 
        -- If older PG, this might fail, but it's standard now.
        INSERT INTO pets (
            id, name, date_of_birth, breed, sex, status, descriptions, profile_settings, photos, species
        ) 
        VALUES (
            gen_random_uuid(), 
            'Allison', 
            '2023-12-08', 
            'Domestic Shorthair', 
            'female', 
            'available',
            '{"spotlight": "Sweet Torti. A total lovebug who thrives as the only pet."}', 
            '{"isSpotlightFeatured": "true"}',
            '[]',
            'Cat'
        );
    END IF;
END $$;

-- 4. Ensure other spotlight pets are set
UPDATE pets
SET profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"')
WHERE name IN ('Sally', 'Merry', 'Colby');

-- DOWN
