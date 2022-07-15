package postgresql

const GetMunicipalitiesByCityId = `
SELECT
	label,
	id,
	slug
FROM municipalities
WHERE city_id = $1
`
