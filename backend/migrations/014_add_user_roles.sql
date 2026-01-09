ALTER TABLE users ADD COLUMN IF NOT EXISTS role text NOT NULL DEFAULT 'tier_1';

ALTER TABLE users ADD CONSTRAINT users_role_check CHECK (role IN ('super_admin', 'admin', 'tier_1', 'tier_2', 'teen'));
