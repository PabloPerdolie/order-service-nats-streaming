CREATE TABLE IF NOT EXISTS delivery (
    delivery_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    zip VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS payment (
    payment_id SERIAL PRIMARY KEY,
    transaction VARCHAR(255) NOT NULL,
    request_id VARCHAR(255) NULL,
    currency VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    payment_dt INT NOT NULL,
    bank VARCHAR(255) NOT NULL,
    delivery_cost INT NOT NULL,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL
);

CREATE TABLE IF NOT EXISTS item (
    item_id SERIAL PRIMARY KEY,
    chrt_id INT NOT NULL UNIQUE,
    track_number VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    rid VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    sale INT NOT NULL,
    size VARCHAR(255) NOT NULL,
    total_price INT NOT NULL,
    nm_id INT NOT NULL,
    brand VARCHAR(255) NOT NULL,
    status INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) NOT NULL UNIQUE,
    track_number VARCHAR(255) NOT NULL,
    entry VARCHAR(255) NOT NULL,
    locale VARCHAR(255) NOT NULL NOT NULL,
    internal_signature VARCHAR(255) NULL,
    customer_id VARCHAR(255) NOT NULL,
    delivery_service VARCHAR(255) NOT NULL,
    shardkey VARCHAR(255) NOT NULL,
    sm_id INT NOT NULL,
    date_created DATE NOT NULL NOT NULL,
    oof_shard VARCHAR(255) NOT NULL,
    delivery_id INT NOT NULL,
    payment_id INT NOT NULL,
    CONSTRAINT fk_delivery_id FOREIGN KEY (delivery_id) REFERENCES delivery(delivery_id),
    CONSTRAINT fk_payment_id FOREIGN KEY (payment_id) REFERENCES payment(payment_id)
);

CREATE INDEX IF NOT EXISTS fk_delivery_id ON orders (delivery_id);

CREATE INDEX IF NOT EXISTS fk_payment_id ON orders (payment_id);

CREATE TABLE IF NOT EXISTS order_item (
    id        SERIAL PRIMARY KEY,
    order_uid text NOT NULL,
    chrt_id   integer NOT NULL,
    CONSTRAINT fk_order_uid FOREIGN KEY(order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE,
    CONSTRAINT fk_chrt_id FOREIGN KEY(chrt_id) REFERENCES item(chrt_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS fk_chrt_id ON order_item (chrt_id);

CREATE INDEX IF NOT EXISTS fk_order_uid ON order_item (order_uid);