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
		and EXTRACT(YEAR FROM date_occurred) = $2
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

https://www.postgresql.org/docs/current/sql-update.html

// WITH exceeded_max_retries AS (
//   SELECT w.ctid FROM work_item AS w
//     WHERE w.status = 'active' AND w.num_retries > 10
//     ORDER BY w.retry_timestamp
//     FOR UPDATE
//     LIMIT 5000
// )
// UPDATE work_item SET status = 'failed'
//   FROM exceeded_max_retries AS emr
//   WHERE work_item.ctid = emr.ctid;

// { userId, year }
func ApplyPreAssignmentsToExpendituresForYear(args []any) error {

	return `
		WITH categorized AS (
			SELECT e.id, e.user_id, bca.category_id, e.description
			from expenditure e 
			JOIN budget_category_preassignment bca
			on bca.description = e.description
			WHERE e.user_id = $1 
			AND e.category_id is null
			AND extract(year from e.date_occurred) = $2
		)
		UPDATE expenditure e
		SET category_id = categorized.category_id
		FROM categorized
		WHERE e.id = categorized.id;
	`
}
