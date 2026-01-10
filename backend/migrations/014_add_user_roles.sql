-- "Nuke and Pave" strategy to fix inconsistent state
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_check;
ALTER TABLE users DROP COLUMN IF EXISTS role;

-- Create fresh
ALTER TABLE users ADD COLUMN role text NOT NULL DEFAULT 'tier_1';

-- Verification constraint
ALTER TABLE users ADD CONSTRAINT users_role_check CHECK (role IN ('super_admin', 'admin', 'tier_1', 'tier_2', 'teen'));
