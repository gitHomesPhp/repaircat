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

CREATE UNIQUE INDEX IF NOT EXISTS u_idx_sc_name_location_id
    ON sc (name, location_id);
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS service (
    id BIGSERIAL PRIMARY KEY,
    label VARCHAR ( 255 )
);

CREATE TABLE IF NOT EXISTS sc_service (
    sc_id BIGINT,
    service_id BIGINT,

    CONSTRAINT fk_sc_service_sc
        FOREIGN KEY (sc_id)
        REFERENCES sc(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_sc_service_service
        FOREIGN KEY (service_id)
        REFERENCES service(id)
        ON DELETE CASCADE
);

ALTER TABLE location
    ALTER COLUMN underground_id SET DEFAULT NULL,
    ADD COLUMN latitude VARCHAR ( 31 ),
    ADD COLUMN longitude VARCHAR ( 31 )
;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
CREATE TABLE IF NOT EXISTS municipalities(
     id BIGSERIAL PRIMARY KEY,
     label VARCHAR ( 63 ) NOT NULL,
     city_id BIGINT,

     CONSTRAINT fk_municipalities_city
         FOREIGN KEY (city_id)
         REFERENCES city(id)
         ON DELETE CASCADE
);

CREATE TYPE region_type AS ENUM ('underground', 'municipality');

CREATE TABLE IF NOT EXISTS location_regions(
    location_id BIGINT,
    region_type region_type,
    region_id BIGINT,

    CONSTRAINT fk_location_regions_location
        FOREIGN KEY (location_id)
        REFERENCES location(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS visitor(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR ( 63 ),
    phone VARCHAR ( 31 ) UNIQUE,
    email VARCHAR ( 255 ) UNIQUE,
    password VARCHAR (127)
);

CREATE TABLE IF NOT EXISTS manager(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR ( 63 ),
    sc_id BIGINT,
    password VARCHAR (127)
);

CREATE TABLE IF NOT EXISTS review(
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    rating SMALLINT NOT NULL,
    visitor_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT fk_review_visitor
        FOREIGN KEY (visitor_id)
        REFERENCES visitor(id)
        ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS review_answer(
     id BIGSERIAL PRIMARY KEY,
     text TEXT NOT NULL,
     manager_id BIGINT NOT NULL,
     created_at TIMESTAMP WITH TIME ZONE,
     updated_at TIMESTAMP WITH TIME ZONE,
     deleted_at TIMESTAMP WITH TIME ZONE,

     CONSTRAINT fk_review_answer_manager
        FOREIGN KEY (manager_id)
        REFERENCES manager(id)
);

INSERT INTO location_regions(
    location_id,
    region_type,
    region_id
)
SELECT
    id as location_id,
    'underground' as region_type,
    underground_id as region_id
FROM location
;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
ALTER TABLE location
    DROP column underground_id
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;