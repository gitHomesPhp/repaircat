package review_external_repository

const CountScReviews = `
	SELECT
    	count(id)
	FROM review_external
	WHERE sc_id = $1
`
