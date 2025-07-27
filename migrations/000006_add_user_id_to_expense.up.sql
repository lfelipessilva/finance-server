-- Add user_id column to expenses table (initially nullable)
ALTER TABLE expenses ADD COLUMN user_id INT;

-- Create a default user if no users exist
INSERT INTO users (id, provider, provider_user_id, email, name, created_at, updated_at)
SELECT 1, 'default', 'default_user', 'default@example.com', 'Default User', NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM users WHERE id = 1);

-- Update existing expenses to have the default user_id
UPDATE expenses SET user_id = 1 WHERE user_id IS NULL;

-- Now make the column NOT NULL
ALTER TABLE expenses ALTER COLUMN user_id SET NOT NULL;

-- Add foreign key constraint
ALTER TABLE expenses ADD CONSTRAINT fk_expenses_user_id 
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- Add index for better query performance
CREATE INDEX idx_expenses_user_id ON expenses (user_id);
