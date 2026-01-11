ALTER TABLE volunteers 
ADD COLUMN IF NOT EXISTS birthday date,
ADD COLUMN IF NOT EXISTS emergency_contact_name text DEFAULT '',
ADD COLUMN IF NOT EXISTS emergency_contact_phone text DEFAULT '',
ADD COLUMN IF NOT EXISTS interest_reason text DEFAULT '',
ADD COLUMN IF NOT EXISTS volunteer_experience text DEFAULT '';
