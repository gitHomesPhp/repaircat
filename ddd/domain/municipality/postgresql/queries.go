package postgresql

const GetMunicipalitiesByCityId = `
SELECT
	label,
	id
FROM municipalities
WHERE city_id = $1
`
