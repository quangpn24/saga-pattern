
-- +migrate Up
CREATE TABLE IF NOT EXISTS orders (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    status varchar(256) NOT NULL,
    customer_id uuid NOT NULL,
    total_amount decimal(14,2),
    note text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS order_item (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity int,
    price decimal(14,2),
    total decimal(14,2),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
    );
ALTER TABLE "order_item"
ADD constraint "fk_order-item"
FOREIGN KEY ("order_id") REFERENCES "orders"("id") ON DELETE CASCADE;
-- +migrate Down
