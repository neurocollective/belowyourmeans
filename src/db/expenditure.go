package db

import (
	"database/sql"
	"log"
	ncsql "github.com/neurocollective/go_utils/sql"
)

// a valid SQLReporter will need pointers as every field. `SQLReporter` implies this.
type Expenditure struct {
	Id           *int     `ncsql:"id",json:"id"`
	UserId       *int     `ncsql:"user_id",json:"userId"`
	CategoryId   *int     `ncsql:"category_id",json:"categoryId"`
	Value        *float32 `ncsql:"value",json:"value"`
	Description  *string  `ncsql:"description",json:"description"`
	DateOccurred *string  `ncsql:"date_occurred",json:"dateOccurred"`
	CreateDate   *string  `ncsql:"create_date",json:"createDate"`
	ModifiedDate *string  `ncsql:"modified_date",json:"modifiedDate"`
}

func (e Expenditure) Zero() Expenditure {

	new := Expenditure{}

	one := 0
	two := 0
	three := 0
	four := float32(0)
	five := ""
	six := ""
	seven := ""
	eight := ""

	new.Id = &one
	new.UserId = &two
	new.CategoryId = &three
	new.Value = &four
	new.Description = &five
	new.DateOccurred = &six
	new.CreateDate = &seven
	new.ModifiedDate = &eight

	return new
}

func ScanExpenditureRow(rows *sql.Rows, receiver *Expenditure) error {

	values := []any{
		&receiver.Id,
		&receiver.UserId,
		&receiver.CategoryId,
		&receiver.Value,
		&receiver.Description,
		&receiver.DateOccurred,
		&receiver.CreateDate,
		&receiver.ModifiedDate,
	}

	err := rows.Scan(values...) 

	if err != nil {
		log.Println("scan error during ScanExpenditureRow(...)")
		return err
	}

	return nil
}

func GetExpenditures(client ncsql.PGClient, query string, args []any) ([]Expenditure, error) {
	rows, err := client.Query(query, args...)

	if err != nil {
		return []Expenditure{}, nil
	}

	// var empty Expenditure
	// empty = empty.Zero()

	// var empty []T

	capacity := 100

	rowArray := make([]Expenditure, capacity, capacity)
	var index int

	for rows.Next() {

		var receiver Expenditure
		// zeroedStruct := receiver.Zero()

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

