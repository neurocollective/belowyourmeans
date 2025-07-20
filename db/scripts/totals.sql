with category_averages as (
        SELECT bc.display_name,
        ROUND(
                (SUM(e.value) / 12), 0
        ) AS monthly_average
        FROM budget_category bc
        JOIN expenditure e ON e.category_id = bc.id
        WHERE extract(year from e.date_occurred) = '2024'
        AND bc.id != 41
        GROUP BY bc.display_name
)
select * from category_averages
UNION
select 'TOTAL', sum(monthly_average) from category_averages
ORDER BY monthly_average;
