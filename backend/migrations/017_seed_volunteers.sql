-- Add unique constraint to email if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'volunteers_email_key'
    ) THEN
        ALTER TABLE volunteers ADD CONSTRAINT volunteers_email_key UNIQUE (email);
    END IF;
END $$;

INSERT INTO volunteers (
    first_name, last_name, email, phone, role, status, bio, photo_url, reliability_score, total_hours, streak, join_date, allergies, skills, position_preferences, availability, badges, address, city, zip
) VALUES
('Linda', 'Taira', 'linda.taira@example.com', '555-0201', 'Tier 2', 'active', 'Dedicated volunteer with a passion for senior dogs.', '', 100, 320, 24, '2022-03-15', false, '{Cat Behavior Training,Medication Administration}', '{Feeding/Cleaning,Cat Socializing}', '{"10AM – 12PM","2PM – 4PM"}', '{New Recruit,Rookie,Regular,Veteran,10 Hours,50 Hours,Bronze Service,Silver Service,Gold Service,Poop Scoop Pro}', '101 Pine St', 'Springfield', '90210'),
('Chris', 'Conner', 'chris.conner@example.com', '555-0202', 'Tier 2', 'active', 'Loves organizing the shelter storage and walking the big dogs.', '', 96, 410, 8, '2021-11-10', true, '{Logistics,Feeding/Cleaning}', '{Feeding/Cleaning,Customer Support}', '{"4PM – 6PM"}', '{Veteran,Platinum Service,Weekend Warrior}', '202 Oak Ave', 'Springfield', '90210'),
('Allison', 'Bruins', 'allison.b@example.com', '555-0203', 'Tier 2', 'active', 'Cat expert and foster parent.', '', 98, 180, 15, '2023-01-05', false, '{Cat Socializing,Photography}', '{Cat Socializing}', '{"6PM – 8PM","Adoption Event: 12pm - 2pm"}', '{Cat Whisperer,Silver Service,Matchmaker}', '303 Maple Dr', 'Springfield', '90210'),
('Katelyn', 'Johnson', 'katelyn.j@example.com', '555-0301', 'Tier 1', 'active', 'New to the area, looking to make friends and help pets.', '', 92, 60, 4, '2023-06-15', false, '{Feeding/Cleaning}', '{Feeding/Cleaning}', '{"10AM – 12PM"}', '{50 Hours,Cleaning Crew}', '404 Elm St', 'Springfield', '90210'),
('Alejandra', '', 'alejandra@example.com', '555-0302', 'Tier 1', 'active', 'Enjoys spending time with the shy cats.', '', 88, 45, 2, '2023-08-01', true, '{Cat Socializing}', '{Cat Socializing}', '{"2PM – 4PM"}', '{10 Hours}', '505 Cedar Ln', 'Springfield', '90210'),
('Arianna', 'Medal', 'arianna.m@example.com', '555-0303', 'Tier 1', 'active', 'Always happy to help with laundry and dishes.', '', 95, 55, 6, '2023-07-20', false, '{Feeding/Cleaning}', '{Feeding/Cleaning}', '{"6PM – 8PM"}', '{Laundry Lord,50 Hours}', '606 Birch Rd', 'Springfield', '90210'),
('Bella', '', 'bella@example.com', '555-0304', 'Tier 1', 'pending', 'Just finished orientation.', '', 0, 0, 0, '2023-12-01', false, '{}', '{Feeding/Cleaning}', '{"Adoption Event: 12pm - 2pm","Adoption Event: 2pm - 4pm"}', '{New Recruit}', '707 Spruce St', 'Springfield', '90210'),
('Brandon', 'Siou', 'brandon.siou@example.com', '555-0305', 'Tier 1', 'active', 'Strong and willing to handle large dogs.', '', 90, 80, 3, '2023-05-10', false, '{Feeding/Cleaning,Maintenance}', '{Feeding/Cleaning}', '{"Adoption Event: 2pm - 4pm"}', '{50 Hours,Dog Walker}', '808 Ash Ave', 'Springfield', '90210'),
('Katie', '', 'katie@example.com', '555-0306', 'Tier 1', 'inactive', 'Moved out of town temporarily.', '', 100, 120, 0, '2022-09-01', false, '{Feeding/Cleaning}', '{Feeding/Cleaning}', '{}', '{Bronze Service}', '909 Walnut St', 'Springfield', '90210'),
('Leanne', '', 'leanne@example.com', '555-0307', 'Tier 1', 'active', 'Loves puppies!', '', 85, 25, 2, '2023-10-05', false, '{Cat Socializing}', '{Cat Socializing}', '{"6PM – 8PM"}', '{10 Hours,Puppy Pal}', '1010 Chestnut Dr', 'Springfield', '90210'),
('Lindsey', '', 'lindsey@example.com', '555-0308', 'Tier 1', 'active', 'Detail oriented cleaner.', '', 97, 95, 10, '2023-04-15', true, '{Feeding/Cleaning,Organization}', '{Feeding/Cleaning}', '{"10AM – 12PM"}', '{Cleaning Crew,50 Hours}', '1111 Willow Way', 'Springfield', '90210'),
('Lynn', '', 'lynn@example.com', '555-0309', 'Tier 1', 'active', 'Retired teacher loves reading to cats.', '', 100, 300, 40, '2023-02-20', false, '{Cat Socializing}', '{Cat Socialization}', '{"10AM – 12PM","12PM – 2PM"}', '{Cat Whisperer,Silver Service,Reliable}', '1212 Magnolia Ct', 'Springfield', '90210'),
('Michelle', 'Black', 'michelle.b@example.com', '555-0310', 'Tier 1', 'active', 'Photographer by trade.', '', 94, 70, 5, '2023-09-10', false, '{Photography,Cat Socializing}', '{Customer Support,Cat Socializing}', '{"Adoption Event: 2pm - 4pm"}', '{Photographer,50 Hours}', '1313 Palm St', 'Springfield', '90210'),
('Nathan', '', 'nathan@example.com', '555-0311', 'Tier 1', 'pending', 'Waiting for training session.', '', 0, 0, 0, '2023-12-15', false, '{}', '{Feeding/Cleaning,Customer Support}', '{"Adoption Event: 12pm - 2pm","Adoption Event: 2pm - 4pm"}', '{New Recruit}', '1414 Sequoia Blvd', 'Springfield', '90210'),
('Sonia', '', 'sonia@example.com', '555-0312', 'Tier 1', 'active', 'Weekend regular.', '', 91, 65, 4, '2023-06-01', false, '{Feeding/Cleaning}', '{Feeding/Cleaning}', '{"10AM – 12PM"}', '{50 Hours,Weekend Warrior}', '1515 Cypress Dr', 'Springfield', '90210'),
('Violet', '', 'violet@example.com', '555-0313', 'Tier 1', 'active', 'Loves the small dogs.', '', 98, 20, 8, '2023-11-01', false, '{Cat Socializing}', '{Cat Socializing}', '{"2PM – 4PM"}', '{10 Hours,Rookie}', '1616 Redwood Ln', 'Springfield', '90210'),
('Evi', '', 'evi@example.com', '555-0401', 'Teen', 'active', 'High school junior.', '', 96, 85, 12, '2023-06-20', false, '{Cat Socializing}', '{Cat Socializing}', '{"4PM – 6PM"}', '{Rising Star,50 Hours}', '1717 Poplar Pl', 'Springfield', '90210')
ON CONFLICT (email) DO NOTHING;
