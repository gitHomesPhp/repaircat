;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS city (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR ( 7 ) UNIQUE NOT NULL,
    label VARCHAR ( 63 ) NOT NULL
);
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS underground (
   id BIGSERIAL PRIMARY KEY,
   label VARCHAR ( 63 ) NOT NULL,
   city_id BIGINT,

   CONSTRAINT fk_underground_city
       FOREIGN KEY (city_id)
           REFERENCES city(id)
           ON DELETE CASCADE
)
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS location (
    id BIGSERIAL PRIMARY KEY,
    city_id BIGINT,
    underground_id BIGINT,
    address VARCHAR ( 511 ),

    CONSTRAINT fk_location_city
        FOREIGN KEY (city_id)
            REFERENCES city(id)
            ON DELETE SET NULL,

    CONSTRAINT fk_location_underground
        FOREIGN KEY (underground_id)
            REFERENCES underground(id)
            ON DELETE SET NULL
);
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS sc (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR ( 63 ) NOT NULL,
    description TEXT DEFAULT '',
    phone VARCHAR ( 31 ) DEFAULT '',
    email VARCHAR ( 255 ) DEFAULT '',
    site VARCHAR ( 1023 ) DEFAULT '',
    location_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT fk_sc_location
        FOREIGN KEY (location_id)
            REFERENCES location(id)
            ON DELETE SET NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS u_idx_sc_name_location_id ON sc (name, location_id)
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
