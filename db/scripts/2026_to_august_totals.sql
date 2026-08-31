with category_averages as (
        SELECT bc.display_name,
        ROUND(
                (SUM(e.value) / 7), 0
        ) AS monthly_average
        FROM budget_category bc
        JOIN expenditure e ON e.category_id = bc.id
        WHERE extract(year from e.date_occurred) = '2026'
        AND e.user_id = 1
        AND bc.id != 41
        GROUP BY bc.display_name
)
select * from category_averages
UNION
select 'TOTAL', sum(monthly_average) from category_averages
ORDER BY monthly_average;

--

select id, description, value, date_occurred from expenditure
where category_id = ?
and extract(year from date_occurred) = '2026';

-- i make about 5700 per pay period, or 11,400 per month, now that 401k is paused
