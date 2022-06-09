CREATE TABLE IF NOT EXISTS sc (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR ( 63 ) UNIQUE NOT NULL,
    description TEXT,
    phone VARCHAR ( 31 ),
    email VARCHAR ( 255 ),
    site VARCHAR ( 1023 ),
    location_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS location (
    id BIGSERIAL PRIMARY KEY,
    city VARCHAR ( 63 ),
    address VARCHAR ( 511 ),
    underground VARCHAR (255)
);