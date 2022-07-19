package postgresql

const SelectScReviewsExternalByCursor = `
SELECT 
	id,
	text,
	rating,
	visitor_id,
	CASE WHEN visitor_id=1 THEN '2GIS'
		 ELSE 'empty'
	END as source,
	'Клиент' as name
FROM review_external
WHERE sc_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`
