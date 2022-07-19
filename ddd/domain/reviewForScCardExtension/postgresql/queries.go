package postgresql

const SelectScReviewsExternalByCursor = `
SELECT 
	id,
	text,
	rating,
	visitor_id
FROM review_external
WHERE sc_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`
