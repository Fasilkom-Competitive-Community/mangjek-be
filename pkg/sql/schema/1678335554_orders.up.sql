CREATE TABLE IF NOT EXISTS orders
(
    id               VARCHAR(150) PRIMARY KEY,
    user_id          VARCHAR(150) REFERENCES users (id)                  NOT NULL,
    driver_id        INT REFERENCES drivers (id)                         NOT NULL,
    order_inquiry_id VARCHAR(150) REFERENCES order_inquiries (id) UNIQUE NOT NULL,
    payment_id       VARCHAR(150) REFERENCES payments (id) UNIQUE        NOT NULL,
    status           VARCHAR(50)                                         NOT NULL,

    created_at       TIMESTAMP                                           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP                                           NOT NULL DEFAULT CURRENT_TIMESTAMP
)