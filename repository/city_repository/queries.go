package city_repository

const InsertCity = `
	INSERT INTO 
		city(label, code) 
		VALUES($1, $2)
`
const SelectCities = `
	SELECT
		id,
		label,
		code
	FROM city
`
