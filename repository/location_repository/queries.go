package location_repository

const InsertLocation = `
	INSERT INTO 
		location(city_id, address, latitude, longitude) 
		VALUES($1, $2, $3, $4)
`

const InsertLocationUndergrounds = `
	INSERT INTO
		location_regions(location_id, region_type, region_id)
		VALUES ($1, 'underground', $2)
`

const GetIdByValues = `
	SELECT
		id
	FROM location
	WHERE 
		city_id = $1 
		AND address = $2 
		AND latitude = $3
		AND longitude = $4
`
