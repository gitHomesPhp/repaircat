package sc_repository

const InsertSc = `
	INSERT INTO 
		sc(name, description, phone, email, site, location_id) 
		VALUES($1, $2, $3, $4, $5, $6)
`
