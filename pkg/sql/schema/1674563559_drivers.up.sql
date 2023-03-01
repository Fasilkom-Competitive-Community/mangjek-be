CREATE TABLE IF NOT EXISTS drivers
(
    id             SERIAL PRIMARY KEY,
    user_id        VARCHAR(150) REFERENCES users (id) UNIQUE NOT NULL,
    police_number  VARCHAR(50)                               NOT NULL,
    vehicle_model  VARCHAR(255)                              NOT NULL,
    vehicle_type   VARCHAR(50)                               NOT NULL,
    nik            VARCHAR(100)                              NOT NULL,
    address        VARCHAR(255)                              NOT NULL,
    is_sim_active  boolean                                   NOT NULL,
    is_stnk_active boolean                                   NOT NULL,

    created_at     TIMESTAMP                                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP                                 NOT NULL DEFAULT CURRENT_TIMESTAMP
);