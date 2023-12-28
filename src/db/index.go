package db

import (
	"database/sql"
	"log"
	"errors"
)

const (
	USER_QUERY = "SELECT id, first_name, last_name, email from budget_user where id = $1;"
	EXPENDITURE_QUERY_STEM = "SELECT id, category_id, value, description, date_occurred from expenditure where user_id = $1"
)

type TestStruct struct {
	id int `nc:"id"`
	name string `nc:"name"`
}

func ScanForTestStruct(rows *sql.Rows, tester *TestStruct) error {

	log.Println("tester inside ScanForTestStruct:", tester)

	if rows == nil {
		return errors.New("rows is nil inside ScanForTestStruct")
	}

	// if tester == nil {
	// 	tester = new(TestStruct)
	// }

	idPointer := &tester.id
	namePointer := &tester.name

	scanError := rows.Scan(idPointer, namePointer)

	log.Println("idPointer inside ScanForTestStruct:", *idPointer)
	log.Println("namePointer inside ScanForTestStruct:", *namePointer)

	if scanError != nil {
		return scanError
	}

	log.Println("tester id:", tester.id)
	log.Println("tester name:", tester.name)

	return nil
}

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
