package db

import (
	"database/sql"
	"errors"
)

const (
	USER_QUERY = "SELECT id, first_name, last_name, email from budget_user where id = $1;"
	EXPENDITURE_QUERY_STEM = "SELECT id, user_id, category_id, value, description, date_occurred from expenditure where user_id = $1"
	CREATE_EXPENDITURE_QUERY_STEM = "insert into expenditure values (nextval('expenditure_id_seq'), $1, $2, $3, $4, $5, now(), now())"
)

type User struct {
	Id int
	FirstName string
	LastName string
	Email string
}

func ScanForUser(rows *sql.Rows, user *User) error {

	if rows == nil {
		return errors.New("rows is nil inside ScanForUser")
	}

	idPointer := &user.Id
	firstNamePointer := &user.FirstName
	lastNamePointer := &user.LastName
	emailPointer := &user.Email

	scanError := rows.Scan(idPointer, firstNamePointer, lastNamePointer, emailPointer)

	if scanError != nil {
		return scanError
	}

	return nil
}

type Expenditure struct {
	Id int
	UserId int
	CategoryId *int
	Value float32
	Description string
	DateOccurred string
}

func GetExpenditureColumnNameByQueryKey(key string) string {
	if key == "amount" {
		return "value"
	} else if key == "user" {
		return "user_id"
	} else if key == "category" {
		return "category_id"
	}
	return ""
}

func ScanForExpenditure(rows *sql.Rows, ex *Expenditure) error {

	if rows == nil {
		return errors.New("rows is nil inside ScanForExpenditure")
	}

	if ex == nil {
		return errors.New("ex is nil inside ScanForExpenditure")		
	}

	idPointer := &ex.Id
	userIdPointer := &ex.UserId
	categoryIdPointer := &ex.CategoryId
	valuePointer := &ex.Value
	descriptionPointer := &ex.Description
	dateOccurredPointer := &ex.DateOccurred

	// send category_id = -1 in the query to ask for `NULL` entries.
	// if ex.CategoryId != nil && *ex.CategoryId == 0 {
	// 	categoryIdPointer = nil
	// }

	scanError := rows.Scan(idPointer, userIdPointer, categoryIdPointer, valuePointer, descriptionPointer, dateOccurredPointer)

	if scanError != nil {
		return scanError
	}

	return nil
}
