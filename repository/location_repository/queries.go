package location_repository

const InsertLocation = `
	INSERT INTO 
		location(city_id, address, underground_id) 
		VALUES($1, $2, $3)
`

const GetIdByValues = `
	SELECT
		id
	FROM location
	WHERE city_id = $1 and address = $2 and underground_id = $3
`
