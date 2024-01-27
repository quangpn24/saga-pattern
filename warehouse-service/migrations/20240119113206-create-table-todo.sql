
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    status varchar(256) NOT NULL,
    content text NOT NULL,
    note text NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
    );
-- +migrate Down
DROP TABLE todos;
