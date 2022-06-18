package location_repository

const InsertLocation = `
	INSERT INTO 
		location(city, address, underground) 
		VALUES($1, $2, $3)
`

const GetIdByValues = `
	SELECT
		id
	FROM location
	WHERE city = $1 and address = $2 and underground = $3
`
