
--+goose Up
--+goose StatementBegin

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER NOT NULL REFERENCES carts (id) ON DELETE CASCADE,
    product VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_cart_id ON cart_items (cart_id);

--+goose StatementEnd

--+goose Down
--+goose StatementBegin

DROP INDEX IF EXISTS idx_cart_id;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;

--+goose StatementEnd
