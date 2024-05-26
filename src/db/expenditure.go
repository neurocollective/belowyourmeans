package db

import (
	"database/sql"
	"log"
	// ncsql "github.com/neurocollective/go_utils/sql"
)

type SQLGenerated interface {
	GetId() sql.NullInt64            // get the id
	Keys() []string         // get the struct pointer names as strings equal to column names, in db column order
	Values() []sql.Null[any]           // get the struct pointer values in db column order
	Get(string) (sql.Null[any], error) // get a struct field by string key - defined by `ncsql:"fieldName"` tag
	TableName() string       // get the table name this struct targets
}

// type Expenditure struct {
// 	Id           sql.NullInt64     `ncsql:"id",json:"id"`
// 	UserId       sql.NullInt64     `ncsql:"user_id",json:"userId"`
// 	CategoryId   sql.NullInt64     `ncsql:"category_id",json:"categoryId"`
// 	Value        sql.NullFloat64   `ncsql:"value",json:"value"`
// 	Description  sql.NullString    `ncsql:"description",json:"description"`
// 	DateOccurred sql.NullTime      `ncsql:"date_occurred",json:"dateOccurred"`
// 	CreateDate   sql.NullTime      `ncsql:"create_date",json:"createDate"`
// 	ModifiedDate sql.NullTime      `ncsql:"modified_date",json:"modifiedDate"`
// }

type Expenditure struct {
	Id           sql.Null[int64]    `ncsql:"id",json:"id"`
	UserId       sql.Null[int64]    `ncsql:"user_id",json:"userId"`
	CategoryId   sql.Null[int64]    `ncsql:"category_id",json:"categoryId"`
	Value        sql.Null[float64]  `ncsql:"value",json:"value"`
	Description  sql.Null[string]   `ncsql:"description",json:"description"`
	DateOccurred sql.Null[string]   `ncsql:"date_occurred",json:"dateOccurred"`
	CreateDate   sql.Null[string]   `ncsql:"create_date",json:"createDate"`
	ModifiedDate sql.Null[string]   `ncsql:"modified_date",json:"modifiedDate"`
}

func ScanExpenditureRow(rows *sql.Rows, expenditure *Expenditure) error {

	values := []any{
		&expenditure.Id,
		&expenditure.UserId,
		&expenditure.CategoryId,
		&expenditure.Value,
		&expenditure.Description,
		&expenditure.DateOccurred,
		&expenditure.CreateDate,
		&expenditure.ModifiedDate,
	}

	err := rows.Scan(values...) 

	if err != nil {
		log.Println("scan error during ScanExpenditureRow(...)")
		return err
	}

	return nil
}

func SelectExpenditure(client ncsql.PGClient, query string, args []any) ([]Expenditure, error) {
	rows, err := client.Query(query, args...)

	if err != nil {
		return []Expenditure{}, nil
	}

	capacity := 100

	rowArray := make([]Expenditure, capacity, capacity)
	var index int

	for rows.Next() {

		var receiver Expenditure

		if index == capacity-1 {
			capacity += 100
			newRowArray := make([]Expenditure, 0, capacity)

			copy(newRowArray, rowArray)
			rowArray = newRowArray
		}

		err := ScanExpenditureRow(rows, &receiver)

		if err != nil {
			log.Println("scanError", err.Error())
			return []Expenditure{}, err
		}

		rowArray[index] = receiver
		index++
	}

	getNextRowError := rows.Err()

	if getNextRowError != nil {
		log.Println("error getting next row:", getNextRowError.Error())
		return []Expenditure{}, getNextRowError
	}

	return rowArray[:index], nil
}

