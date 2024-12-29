package queries

// this query needs a review, probably more complicated than necessary
func SelectExpendituresWithCategoryNameByUserAndMonth() string {
	return `
		select e.*, bc.display_name as category_name from expenditure e
		JOIN budget_category bc
		ON e.category_id = bc.id
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and e.value > 0
	UNION
	select *, 'uncategorized' as category_name from expenditure e
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and e.value > 0
		and e.category_id IS NULL;
	`
}

func SelectExpendituresWithCategoryNameByUserUnique() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `select distinct on (description) id, description from expenditure
		where user_id = $1
		and category_id IS NULL
		and value < 0
		GROUP BY description, id;`
}

