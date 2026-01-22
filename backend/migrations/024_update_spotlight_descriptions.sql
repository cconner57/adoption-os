-- UP
DO $$
DECLARE
BEGIN
    -- 1. Ophelia
    IF NOT EXISTS (SELECT 1 FROM pets WHERE name = 'Ophelia') THEN
        INSERT INTO pets (id, name, species, sex, status, descriptions, profile_settings, photos)
        VALUES (
            gen_random_uuid(), 'Ophelia', 'Cat', 'female', 'available',
            '{"spotlight": "Chatty solo queen! I love pets on my terms and will be your loyal bestie if I choose you."}',
            '{"isSpotlightFeatured": "true"}',
            '[]'
        );
    ELSE
        UPDATE pets SET 
            descriptions = jsonb_set(COALESCE(descriptions, '{}'), '{spotlight}', '"Chatty solo queen! I love pets on my terms and will be your loyal bestie if I choose you."'),
            profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"'),
            status = 'available'
        WHERE name = 'Ophelia';
    END IF;

    -- 2. Tom
    IF NOT EXISTS (SELECT 1 FROM pets WHERE name = 'Tom') THEN
        INSERT INTO pets (id, name, species, sex, status, descriptions, profile_settings, photos)
        VALUES (
            gen_random_uuid(), 'Tom', 'Cat', 'male', 'available',
            '{"spotlight": "Playful & cuddly! Jerry''s brother, looking for fun and snuggles. Very friendly young spirit."}',
            '{"isSpotlightFeatured": "true"}',
            '[]'
        );
    ELSE
        UPDATE pets SET 
            descriptions = jsonb_set(COALESCE(descriptions, '{}'), '{spotlight}', '"Playful & cuddly! Jerry''s brother, looking for fun and snuggles. Very friendly young spirit."'),
            profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"'),
            status = 'available'
        WHERE name = 'Tom';
    END IF;

    -- 3. Colby
    -- Using simple UPDATE; assuming he exists or we can insert if we want to be safe. But user implies he exists.
    -- We'll just UPDATE.
    UPDATE pets SET 
        descriptions = jsonb_set(COALESCE(descriptions, '{}'), '{spotlight}', '"Non-stop purrs & biscuits! Social, clumsy, and very affectionate. I''ll talk your ear off!"'),
        profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"'),
        status = 'available'
    WHERE name = 'Colby';

    -- 4. Merry
    UPDATE pets SET 
        descriptions = jsonb_set(COALESCE(descriptions, '{}'), '{spotlight}', '"Lord of the Strings! Last of my litter, I love to climb, eat, and play. Friendly adventurer!"'),
        profile_settings = jsonb_set(COALESCE(profile_settings, '{}'), '{isSpotlightFeatured}', '"true"'),
        status = 'available'
    WHERE name = 'Merry';

END $$;

-- DOWN
