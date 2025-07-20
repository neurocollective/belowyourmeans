package sqlv2

import (
	"database/sql"
	//"errors"
	"log"
	//"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type PGClient interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type SQLDescriber interface {
	Columns() []any
	ColumnsString() string
	TableName() string
}

func JSONifyNull[T any](column sql.Null[T]) T {
	return column.V
}

type Expenditure struct {
	Id           *int     `ncsql:"id",json:"id"`
	UserId       *int     `ncsql:"user_id",json:"userId"`
	// CategoryId   *int     `ncsql:"category_id",json:"categoryId"`
	CategoryId   *sql.Null[int]     `ncsql:"category_id",json:"categoryId"`
	Value        *float32 `ncsql:"value",json:"value"`
	Description  *string  `ncsql:"description",json:"description"`
	// DateOccurred *string  `ncsql:"date_occurred",json:"dateOccurred"`
	// CreateDate   *string  `ncsql:"create_date",json:"createDate"`
	// ModifiedDate *string  `ncsql:"modified_date",json:"modifiedDate"`
	DateOccurred *sql.Null[string]  `ncsql:"date_occurred",json:"dateOccurred"`
	CreateDate   *sql.Null[string]  `ncsql:"create_date",json:"createDate"`
	ModifiedDate *sql.Null[string]  `ncsql:"modified_date",json:"modifiedDate"`
}

type ExpenditureForJSON struct {
	Id           int     `json:"id"`
	UserId       int     `json:"userId"`
	CategoryId   int     `json:"categoryId"`
	Value        float32 `json:"value"`
	Description  string  `json:"description"`
	DateOccurred string  `json:"dateOccurred"`
	CreateDate   string  `json:"createDate"`
	ModifiedDate string  `json:"modifiedDate"`
}

func (e Expenditure) ForJson() ExpenditureForJSON {
	var forJSON ExpenditureForJSON

	forJSON.Id = *e.Id
	forJSON.UserId = *e.UserId
	forJSON.CategoryId = JSONifyNull[int](*e.CategoryId)
	forJSON.Value = *e.Value
	forJSON.Description = *e.Description
	forJSON.DateOccurred = JSONifyNull[string](*e.DateOccurred)
	forJSON.CreateDate = JSONifyNull[string](*e.CreateDate)
	forJSON.ModifiedDate = JSONifyNull[string](*e.ModifiedDate)

	return forJSON
}

func (e Expenditure) ColumnsString() string {
	columns := []string{
		"exp.id",
		"exp.user_id",
		"exp.category_id",
		"exp.value",
		"exp.description",
		"exp.date_occurred",
		"exp.create_date",
		"exp.modified_date",
	}
	return strings.Join(columns, ",")
}

func (e Expenditure) Columns() []any {
	return []any{
		e.Id,
		e.UserId,
		e.CategoryId,
		e.Value,
		e.Description,
		e.DateOccurred,
		e.CreateDate,
		e.ModifiedDate,
	}
}

func (e *Expenditure) Zero() {
	id := 0
	e.Id = &id
	userId := 0
	e.UserId = &userId

	// categoryId := 0
	// e.CategoryId = &categoryId

	categoryId := new(sql.Null[int])
	e.CategoryId = categoryId

	var value float32 = 0
	e.Value = &value
	description := ""
	e.Description = &description

	dateOccurred := new(sql.Null[string])
	e.DateOccurred = dateOccurred
	createDate := new(sql.Null[string])
	e.CreateDate = createDate
	modifiedDate := new(sql.Null[string])
	e.ModifiedDate = modifiedDate

	// dateOccurred := ""
	// e.DateOccurred = &dateOccurred
	// createDate := ""
	// e.CreateDate = &createDate
	// modifiedDate := ""
	// e.ModifiedDate = &modifiedDate
}

func GetZeroedExpenditures(size int) ([]*Expenditure) {
	expenditures := make([]*Expenditure, size, size)

	for index, _ := range expenditures {
		e := new(Expenditure)
		e.Zero()
		expenditures[index] = e
	}
	return expenditures
}

func (e Expenditure) TableName() string {
	return "expenditure exp"
}

// connectionString -> "user=postgres password=postgres dbname=postgres sslmode=disable"
func BuildPostgresClient(connectionString string) (PGClient, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {

		log.Println("ERROR opening postgres connection with github.com/neurocollective/go_utils.BuildPostgresClient() ->")
		log.Println(err.Error())

		return nil, err
	}

	return db, nil
}

func SelectExpenditures(client PGClient, userId int, month int) ([]*Expenditure, error) {

	var empty []*Expenditure

	query := SelectExpendituresWithCategoryNameByUserAndMonth()

	log.Println("query", query)

	args := []any{ userId, month }

	rows, err := client.Query(query, args...)

	if err != nil {
		return empty, err
	}

	//capacity := 10000

	expenditures := GetZeroedExpenditures(10000)

	// // TODO - type assertion needed?
	// asDescribers, ok := expenditureList.(ArrayList[SQLDescriber])

	// if !ok {
	// 	return empty, errors.New("could not type assert Expenditure to SQLDescriber")
	// }

	correctSizedSlice,err := ReceiveRows[*Expenditure](rows, expenditures)

	if err != nil {
		return empty, err
	}

	return correctSizedSlice, nil
}

func ReceiveRows[T SQLDescriber](rows *sql.Rows, receiverObjects []T) ([]T, error) {

	var empty []T
	var index int

	for rows.Next() {

		var receiver SQLDescriber = receiverObjects[index]

		err := ScanRow(rows, receiver)

		if err != nil {
			log.Println("scanError", err.Error())
			return empty, err
		}

		index++
	}

	err := rows.Err()

	if err != nil {
		log.Println("error getting next row:", err.Error())
		return empty, err
	}

	return receiverObjects[:index], nil
}

func ScanRow(rows *sql.Rows, object SQLDescriber) error {

	values := object.Columns()

	err := rows.Scan(values...)

	if err != nil {
		log.Println("scan error during ScanRow(rows, object) ->")
		log.Println(err)
		return err
	}

	return nil
}

func FillTemplatedQuery(template string, object SQLDescriber) string {
	withColumns := strings.ReplaceAll(template, "$COLUMNS", object.ColumnsString())
	return strings.ReplaceAll(withColumns, "$TABLE_NAME", object.TableName())	
}

const selectExpendituresWithCategoryNameByUserAndMonthQueryTemplate = `
	SELECT $COLUMNS from $TABLE_NAME
		JOIN budget_category bc
		ON exp.category_id = bc.id
		where exp.user_id = $1
		and EXTRACT(MONTH FROM exp.date_occurred) - 1 = $2
		and exp.value > 0
	UNION
	SELECT $COLUMNS from $TABLE_NAME
		where exp.user_id = $1
		and EXTRACT(MONTH FROM exp.date_occurred) - 1 = $2
		and exp.value > 0
		and exp.category_id IS NULL;
	`

// this query needs a review, probably more complicated than necessary
func SelectExpendituresWithCategoryNameByUserAndMonth() string {
	empty := new(Expenditure)
	template := selectExpendituresWithCategoryNameByUserAndMonthQueryTemplate
	return FillTemplatedQuery(template, empty)
}
