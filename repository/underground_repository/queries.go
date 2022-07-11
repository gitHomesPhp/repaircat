package underground_repository

const SelectUndergroundByCityId = `
	SELECT 
		id,
		label
	FROM underground
	WHERE city_id = $1
`
