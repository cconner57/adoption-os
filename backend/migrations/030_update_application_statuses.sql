-- Up Migration
ALTER TABLE applications ALTER COLUMN status SET DEFAULT 'submitted';

UPDATE applications SET status = 'submitted' WHERE status = 'pending';
UPDATE applications SET status = 'adopted' WHERE status = 'approved';
UPDATE applications SET status = 'rejected' WHERE status = 'denied';
UPDATE applications SET status = 'under_review' WHERE status = 'needs_info';

-- Add check constraint to ensure only valid statuses are used (optional but good practice, though we might skip strict constraint to allow for easier future updates if Go handles validation. Let's stick to update and default for now as postgres enums aren't used here, it's text).
