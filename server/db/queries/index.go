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

func SelectExpendituresMissingCategoryNameByUserUnique() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `select distinct on (description) id, description from expenditure
		where user_id = $1
		and category_id IS NULL
		and value > 0
		GROUP BY description, id;`
}


func SelectMonthlyAverageExpendituresByCategoryByUser() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `select round(sum(value) / 12, 2) as total, category_id, bc.display_name as name from expenditure e
		join budget_category bc
		on bc.id = e.category_id
		where e.user_id = 1
		and value > 0
		GROUP BY e.category_id, name;`
}

