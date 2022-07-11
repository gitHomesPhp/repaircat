package postgresql

const SelectScCardExtensionById = `
SELECT
	sc.id,
	COALESCE(name, '') as name,
	COALESCE(description, '') as description,
	COALESCE(phone, '') as phone,
	COALESCE(email, '') as email,
	COALESCE(site, '') as site,
	COALESCE(city.label, '') as city,
	COALESCE(address, '') as address,
	COALESCE(latitude, '') as latitude,
	COALESCE(longitude, '') as longitude,

	location_id
FROM sc LEFT JOIN location ON location_id = location.id
	JOIN city on location.city_id = city.id
WHERE sc.id = $1
`

const GetUndergrounds = `
SELECT
	underground.label
FROM underground LEFT JOIN location_regions ON region_id = underground.id AND location_regions.region_type = 'underground'
WHERE location_id = $1
`

const GetReviewInfo = `
SELECT
    count(sc_id) as count,
    COALESCE(avg(rating), 0) as rating
FROM sc LEFT JOIN review_external ON sc.id = review_external.sc_id
WHERE sc.id = ($1)
GROUP BY sc.id
`
