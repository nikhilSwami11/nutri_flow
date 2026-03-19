CREATE TABLE IF NOT EXISTS pantry_items (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL,
    name        VARCHAR(255) NOT NULL,
    quantity    DECIMAL(10,2) NOT NULL DEFAULT 0,
    unit        VARCHAR(50) NOT NULL DEFAULT 'pieces',
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_pantry_items_user_id ON pantry_items(user_id);