how to get month as integer from timstamp columns in postgres ->

```
SELECT EXTRACT(MONTH FROM date_occurred) FROM expenditure;
```
