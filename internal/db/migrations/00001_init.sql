
--+goose Up
--+goose StatementBegin

CREATE TABLE carts(
                      id BIGSERIAL PRIMARY KEY,
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                      updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE cart_items (
                            id BIGSERIAL PRIMARY KEY,
                            cart_id BIGINT NOT NULL, FOREIGN KEY(cart_id) REFERENCES carts(id) ON DELETE CASCADE,
                            product VARCHAR (255) NOT NULL,
                            price DECIMAL(10,2) NOT NULL CHECK(price>0),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_cart_id ON cart_items(cart_id);

--+goose StatementEnd

--+goose Down
--+goose StatementBegin

DROP INDEX IF EXISTS idx_cart_id;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;

--+goose StatementEnd
