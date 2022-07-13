package postgresql

const GetUndergroundsByCity = `
SELECT
	label,
	id
FROM underground
WHERE city_id = $1
`
