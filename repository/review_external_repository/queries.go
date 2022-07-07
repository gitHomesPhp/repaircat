package review_external_repository

const InsertReviewExternal = `
	INSERT INTO 
		review_external(text, rating, visitor_id, sc_id) 
		VALUES($1, $2, $3, $4)
`
