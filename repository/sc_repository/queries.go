package sc_repository

const SelectId = `
	SELECT id 
	FROM sc 
	where id = $1
`

const SelectScById = `
	SELECT
    	sc.id,
    	COALESCE(name, '') as name,
    	COALESCE(description, '') as description,
    	COALESCE(phone, '') as phone,
    	COALESCE(email, '') as email,
    	COALESCE(site, '') as site,
    	COALESCE(location.id, 0) as location_id,
    	COALESCE(city.label, '') as city,
    	COALESCE(address, '') as address,
    	COALESCE(underground.label, '') as underground,
		COALESCE(latitude, '') as latitude,
		COALESCE(longitude, '') as longitude
	FROM sc LEFT JOIN location ON location_id = location.id
		JOIN city on location.city_id = city.id
		JOIN underground on location.underground_id = underground.id
	WHERE sc.id = $1
`

const InsertSc = `
	INSERT INTO 
		sc(name, description, phone, email, site, location_id) 
		VALUES($1, $2, $3, $4, $5, $6)
`

const SelectScWithLocationPageCursor = `
SELECT
    sc.id,
    COALESCE(name, '') as name,
    COALESCE(description, '') as description,
    COALESCE(phone, '') as phone,
    COALESCE(email, '') as email,
    COALESCE(site, '') as site,
    COALESCE(location.id, 0) as location_id,
    COALESCE(city.label, '') as city,
    COALESCE(address, '') as address,
    COALESCE(underground.label, '') as underground
FROM sc LEFT JOIN location ON location_id = location.id
	JOIN city on location.city_id = city.id
	LEFT JOIN underground on location.underground_id = underground.id
ORDER BY sc.id ASC
LIMIT $1
OFFSET $2
`
