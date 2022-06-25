package underground_repository

const SelectUndergroundByCityId = `
	SELECT 
		id,
		label
	FROM underground
	WHERE city_id = $1
`

const SelectUndergroundsOfLocations = `
	SELECT
	    underground.id,
	    underground.label,
	    location_id
	FROM
	    underground
	        JOIN location_regions ON region_id = underground.id AND location_regions.region_type = 'underground'
	WHERE location_id = ANY ($1)
`
