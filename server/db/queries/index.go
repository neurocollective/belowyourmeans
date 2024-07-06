package queries

func SelectExpendituresWithCategoryNameByUserAndMonth() string {
	return `
		select e.*, bc.display_name as category_name from expenditure e
		JOIN budget_category bc
		ON e.category_id = bc.id
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and e.value < 0
	UNION
	select *, 'uncategorized' as category_name from expenditure e
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and e.value < 0
		and e.category_id IS NULL;
	`
}
