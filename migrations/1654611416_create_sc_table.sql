CREATE TABLE IF NOT EXISTS public.sc (
    sc_id BIGSERIAL PRIMARY KEY,
    name VARCHAR( 63 ) UNIQUE NOT NULL,
    description text,
    phone VARCHAR( 31 ) UNIQUE NULL,
    site VARCHAR ( 1024 ) NULL
);