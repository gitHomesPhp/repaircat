package postgresql

const GetUndergroundsByCity = `
SELECT
	label,
	id,
	slug
FROM underground
WHERE city_id = $1
`
