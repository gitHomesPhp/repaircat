package postgres

const SelectScCardExtensionById = `
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
	COALESCE(latitude, '') as latitude,
	COALESCE(longitude, '') as longitude
FROM sc LEFT JOIN location ON location_id = location.id
	JOIN city on location.city_id = city.id
WHERE sc.id = $1
`
