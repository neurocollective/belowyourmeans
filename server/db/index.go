package db

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"errors"
	"log"
	"io"
	"bytes"

	bymsql "neurocollective.io/neurocollective/belowyourmeans/server/structs/sql"
)

const (
	USER_QUERY = "SELECT id, first_name, last_name, email from budget_user where id = $1;"
	EXPENDITURE_QUERY_STEM = "SELECT id, user_id, category_id, value, description, date_occurred from expenditure where user_id = $1"
	CREATE_EXPENDITURE_QUERY_STEM = "insert into expenditure values (nextval('expenditure_id_seq'), $1, $2, $3, $4, $5, now(), now())"
	CHECK_LOGIN_QUERY = "SELECT * from budget_user where email = $1;"
	CREATE_USER_QUERY = "INSERT INTO budget_user VALUES (nextval('budget_user_id_seq'), $1, $2, $3, $4, now(), now()) RETURNING id;"
)

func ScanForUser(rows *sql.Rows, user *bymsql.User) error {

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

// type Expenditure struct {
// 	Id int
// 	UserId int
// 	CategoryId *int
// 	Value float32
// 	Description string
// 	DateOccurred string
// }

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

// func ScanForExpenditure(rows *sql.Rows, ex *Expenditure) error {

// 	if rows == nil {
// 		return errors.New("rows is nil inside ScanForExpenditure")
// 	}

// 	if ex == nil {
// 		return errors.New("ex is nil inside ScanForExpenditure")		
// 	}

// 	idPointer := &ex.Id
// 	userIdPointer := &ex.UserId
// 	categoryIdPointer := &ex.CategoryId
// 	valuePointer := &ex.Value
// 	descriptionPointer := &ex.Description
// 	dateOccurredPointer := &ex.DateOccurred

// 	// send category_id = -1 in the query to ask for `NULL` entries.
// 	// if ex.CategoryId != nil && *ex.CategoryId == 0 {
// 	// 	categoryIdPointer = nil
// 	// }

// 	scanError := rows.Scan(idPointer, userIdPointer, categoryIdPointer, valuePointer, descriptionPointer, dateOccurredPointer)

// 	if scanError != nil {
// 		return scanError
// 	}

// 	return nil
// }

func ScanForUserLoginData(rows *sql.Rows, user *bymsql.User) error {

	if rows == nil {
		return errors.New("rows is nil inside ScanForUserLoginData")
	}

	if user == nil {
		return errors.New("user is nil inside ScanForUserLoginData")	
	}

	idPointer := &user.Id
	emailPointer := &user.Email
	passwordPointer := &user.HashedPassword

	scanError := rows.Scan(idPointer, emailPointer, passwordPointer)

	if scanError != nil {
		return scanError
	}

	return nil
}

func ScanForUserSignupData(rows *sql.Rows, user *bymsql.User) error {

	if rows == nil {
		return errors.New("rows is nil inside ScanForUserLoginData")
	}

	if user == nil {
		return errors.New("user is nil inside ScanForUserLoginData")	
	}

	idPointer := &user.Id

	scanError := rows.Scan(idPointer)

	if scanError != nil {
		return scanError
	}

	return nil
}

// type NodeQueryExecutor func[T any](query string, queryParams []any, specialRule string) ([]T, error)

func ExecuteNodeQuery[T any](query string, queryParams []any, specialRule string) ([]T, error) {

	client := new(http.Client)

	body := map[string]any{
		"query":      query,
		"parameters": queryParams,
		// "specialRule": "mapExpenditures",
	}

	if specialRule != "" {
		body["specialRule"] = specialRule
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bodyBytes)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:3001/query", bodyReader)

	if err != nil {
		log.Println("creating node request FAILED")
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		log.Println("request to node server FAILED")
		return nil, err
	}

	results := []T{}

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("reading node query response body FAILED")
		return nil, err
	}

	err = json.Unmarshal(bytes, &results)

	if err != nil {
		log.Println("unmarshaling node query response body FAILED, response payload was:")
		log.Println(string(bytes))
		return nil, err
	}

	return results, nil
}

// args -> []any{categoryId, userId, description} -> `int,int,string`
func ApplyBudgetCategoryItems(args []any) error {

	query := "update expenditure set category_id = $1, modified_date = now() where user_id = $2 and description = $3;"
	_, err := ExecuteNodeQuery[any](query, args, "")

	if err != nil {
		return err
	}
	return nil
}

