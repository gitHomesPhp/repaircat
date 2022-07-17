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

const FindLocationIds = `
SELECT
	id
FROM underground
	JOIN location_regions ON underground.id = location_regions.region_id and region_type = 'underground'
WHERE label ILIKE $1 || '%'
`

const FindLocationIdsByUndergroundSlug = `
SELECT
	location_id
FROM underground
	JOIN location_regions ON underground.id = location_regions.region_id and region_type = 'underground'
	JOIN city ON city.id = underground.city_id
WHERE underground.slug = $1 AND city.code = $2
`

const FindLocationIdsByMunicipalitySlug = `
SELECT
	location_id
FROM municipality
	JOIN location_regions ON municipality.id = location_regions.region_id and region_type = 'municipality'
	JOIN city ON city.id = municipality.city_id
WHERE municipality.slug = $1 AND city.code = $2
`

const GetScCardListByLocationIds = `
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
WHERE location.id = ANY ($1)
ORDER BY sc.id
LIMIT $2
OFFSET $3
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

const GetMunicipality = `
SELECT
	municipalities.label,

	location_id
FROM municipalities
	JOIN location_regions ON region_id = municipalities.id AND location_regions.region_type = 'municipality'
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

const PreviousNextQueryByUnderground = `
SELECT
	EXISTS (
		SELECT
			sc.id,
		FROM sc
			JOIN location ON sc.location_id = location.id
			JOIN city ON location.city_id = city.id
		WHERE location.id = ANY ($1)
		ORDER BY sc.id
		LIMIT 1
		OFFSET $2
	) as previous,
	EXISTS (
		SELECT
			sc.id,
		FROM sc
			JOIN location ON sc.location_id = location.id
			JOIN city ON location.city_id = city.id
		WHERE location.id = ANY ($1)
		ORDER BY sc.id
		LIMIT 1
		OFFSET $3
	) as next
`
