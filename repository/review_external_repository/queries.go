package review_external_repository

const InsertReviewExternal = `
	INSERT INTO 
		review_external(text, rating, visitor_id, sc_id) 
		VALUES($1, $2, $3, $4)
`

const CountScReviews = `
	SELECT
    	count(id)
	FROM review_external
	WHERE sc_id = $1
`
