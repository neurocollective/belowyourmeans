package queries

// this query needs a review, probably more complicated than necessary
func SelectExpendituresWithCategoryNameByUserAndMonth() string {
	return `
		select e.*, bc.display_name as category_name from expenditure e
		JOIN budget_category bc
		ON e.category_id = bc.id
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and EXTRACT(YEAR FROM e.date_occurred) = $3
		and e.category_id != 41
		and e.value > 0
		and e.category_id IS NOT NULL
	UNION
	select *, 'uncategorized' as category_name from expenditure e
		where e.user_id = $1
		and EXTRACT(MONTH FROM e.date_occurred) - 1 = $2
		and EXTRACT(YEAR FROM e.date_occurred) = $3
		and e.value > 0
		and e.category_id IS NULL;
	`
}

func SelectExpendituresMissingCategoryNameByUserUnique() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `
		select distinct on (description) id, description from expenditure
		where user_id = $1
		and EXTRACT(YEAR FROM e.date_occurred) = $2
		and category_id IS NULL
		and value > 0
		GROUP BY description, id;
	`
}


func SelectMonthlyAverageExpendituresByCategoryByUser() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `
		select round(sum(value) / 12, 2) as total, category_id, bc.display_name as name
		from expenditure e
		join budget_category bc
		on bc.id = e.category_id
		where e.user_id = $1
		and EXTRACT(YEAR FROM e.date_occurred) = $2
		and value > 0
		and e.category_id != 25
		GROUP BY e.category_id, name;
	`
}

func SelectMonthlyAverageExpendituresByUser() string {
	// "select DISTINCT ON (description), id, description, category_id from expenditure where user_id = $1;"
	//"select description from expenditure where user_id = $1 and category_id IS NULL GROUP BY description HAVING count(description) > 1;"
	return `
		select round(sum(value) / 12, 2) as total
		from expenditure e
		where e.user_id = $1
		and EXTRACT(YEAR FROM e.date_occurred) = $2
		and value > 0
		and e.category_id != 25;
	`
}

// pass arguments:
// userId as int, year as string
func GetAnnualReportQuery() string {
	return `
		with category_averages as (
	        SELECT bc.display_name,
	        ROUND(
	                (SUM(e.value) / 12), 0
	        ) AS monthly_average
	        FROM budget_category bc
	        JOIN expenditure e ON e.category_id = bc.id
	        WHERE extract(year from e.date_occurred) = $1
	        AND bc.id != 41
	        AND bc.user_id = $2
	        GROUP BY bc.display_name
		)
		select * from category_averages
		UNION
		select 'TOTAL', sum(monthly_average) from category_averages
		ORDER BY monthly_average;
`
}

