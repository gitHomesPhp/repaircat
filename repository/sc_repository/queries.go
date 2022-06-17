package sc_repository

const SelectScPage = `
	SELECT
		sc.id,
		COALESCE(name, ''),
		COALESCE(description, ''),
		COALESCE(phone, ''),
		COALESCE(email, ''),
		COALESCE(site, '')
	FROM sc
	WHERE sc.id >= $1 AND sc.id <= $2
`

const SelectId = `
	SELECT id 
	FROM sc 
	where id = $1
`

const InsertSc = `
	INSERT INTO 
		sc(name, description, phone, email, site) 
		VALUES($1, $2, $3, $4, $5)
`

const SelectScPage2 = `
	SELECT
		sc.id,
		COALESCE(name, ''),
		COALESCE(description, ''),
		COALESCE(phone, ''),
		COALESCE(email, ''),
		COALESCE(site, ''),
		COALESCE(city, ''),
		COALESCE(address, ''),
		COALESCE(underground, '')
	FROM sc LEFT JOIN location ON location_id = location.id
	WHERE sc.id >= $1 AND sc.id <= $2
`
