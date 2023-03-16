CREATE TABLE IF NOT EXISTS order_inquiries
(
    id                  VARCHAR(150) PRIMARY KEY,
    user_id             VARCHAR(150) REFERENCES users (id) NOT NULL,
    origin_lat          DOUBLE PRECISION                   NOT NULL,
    origin_long         DOUBLE PRECISION                   NOT NULL,
    origin_address      TEXT                               NOT NULL,
    destination_lat     DOUBLE PRECISION                   NOT NULL,
    destination_long    DOUBLE PRECISION                   NOT NULL,
    destination_address TEXT                               NOT NULL,
    price               BIGINT                             NOT NULL,
    distance            INT                                NOT NULL,
    duration            INT                                NOT NULL,
    routes              TEXT                               NOT NULL,

    created_at          TIMESTAMP                          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP                          NOT NULL DEFAULT CURRENT_TIMESTAMP
)