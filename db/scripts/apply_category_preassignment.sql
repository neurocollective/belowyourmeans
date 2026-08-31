WITH categorized AS (
	SELECT e.id, e.user_id, bca.category_id, e.description
	from expenditure e 
	JOIN budget_category_preassignment bca
	on bca.description = e.description
	WHERE e.user_id = 1 
	AND e.category_id is null
	AND extract(year from e.date_occurred) = '2025'
)
UPDATE expenditure e
SET category_id = categorized.category_id
FROM categorized
WHERE e.id = categorized.id;