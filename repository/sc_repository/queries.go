package sc_repository

const SelectScPage = `
	SELECT
		sc.id,
		COALESCE(name, ''),
		COALESCE(description, ''),
		COALESCE(phone, ''),
		COALESCE(email, ''),
		COALESCE(site, '')
	FROM sc
	WHERE sc.id >= $1 AND sc.id <= $2
`

const SelectId = `
	SELECT id 
	FROM sc 
	where id = $1
`

const InsertSc = `
	INSERT INTO 
		sc(name, description, phone, email, site, location_id) 
		VALUES($1, $2, $3, $4, $5, $6)
`

const SelectScWithLocationPage = `
	SELECT
		sc.id,
		COALESCE(name, ''),
		COALESCE(description, ''),
		COALESCE(phone, ''),
		COALESCE(email, ''),
		COALESCE(site, ''),
		COALESCE(location.id, 0),
		COALESCE(city, ''),
		COALESCE(address, ''),
		COALESCE(underground, '')
	FROM sc LEFT JOIN location ON location_id = location.id
	WHERE sc.id >= $1 AND sc.id <= $2
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
	JOIN underground on location.underground_id = underground.id
ORDER BY sc.id ASC
LIMIT $1
OFFSET $2
`
