package postgresql

const SelectScCardsByCursor = `
SELECT
	sc.id,
	COALESCE(name, '') as name,
    COALESCE(description, '') as description,
    COALESCE(phone, '') as phone,
    COALESCE(email, '') as email,
    COALESCE(site, '') as site,
	COALESCE(address, '') as address,
	COALESCE(latitude, '') as latitude,
	COALESCE(longitude, '') as longitude,
	COALESCE(city.label, '') as city,

	location_id
FROM sc 
	JOIN location ON sc.location_id = location.id
	JOIN city ON location.city_id = city.id
ORDER BY sc.id
LIMIT $1
OFFSET $2
`

const GetScCardInfo = `
SELECT
    sc.id,
    count(sc_id) as count,
    COALESCE(avg(rating), 0) as rating
FROM sc LEFT JOIN review_external ON sc.id = review_external.sc_id
WHERE sc.id = ANY ($1)
GROUP BY sc.id
`

const GetUndergrounds = `
	SELECT
	    underground.label,
	    location_id
	FROM
	    underground
	        JOIN location_regions ON region_id = underground.id AND location_regions.region_type = 'underground'
	WHERE location_id = ANY ($1)
`

const PreviousNextQuery = `
SELECT
	exists(
		select id from sc where id < $1
	) as previous,
	exists(
		select id from sc where id > $2
	) as next
`
