
-- +migrate Up
CREATE TABLE IF NOT EXISTS "customers" (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" varchar(255) NOT NULL,
    balance DECIMAL(14,2) NOT NULL DEFAULT 0,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "transactions" (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id uuid NOT NULL,
    order_id uuid NOT NULL,
    amount DECIMAL(14,2) NOT NULL DEFAULT 0,
    trans_type varchar(50),
    content text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "transactions"
    ADD CONSTRAINT "fk_customer_id"
        FOREIGN KEY ("customer_id") REFERENCES "customers"("id");
-- +migrate Down
