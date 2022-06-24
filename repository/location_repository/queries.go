package location_repository

const InsertLocation = `
	INSERT INTO 
		location(city_id, address, underground_id) 
		VALUES($1, $2, NULLIF($3, 0))
`

const GetIdByValues = `
	SELECT
		id
	FROM location
	WHERE 
		city_id = $1 
		AND address = $2 
		AND underground_id = $3
`
const SelectIdByValuesAndNullUnderground = `
	SELECT
		id
	FROM location
	WHERE 
		city_id = $1 
		AND address = $2 
		AND underground_id is NULL
`
