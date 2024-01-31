
-- +migrate Up
CREATE TABLE IF NOT EXISTS "products" (
    "id" uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" varchar(255) NOT NULL,
    "price" DECIMAL(14,2) NOT NULL,
    "quantity" INTEGER NOT NULL DEFAULT 0,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
