-- Up Migration
ALTER TABLE contracts 
ADD COLUMN IF NOT EXISTS signature text,
ADD COLUMN IF NOT EXISTS signed_at timestamp(0) with time zone;
