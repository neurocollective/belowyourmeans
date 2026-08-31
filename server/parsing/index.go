package parsing

import (
	//"os"
	//"fmt"
	"errors"
	"strings"
	"strconv"
	"log"
	ncsql "github.com/neurocollective/go_utils/sql"
)

const (
	CAPONE_ACCOUNT_NUMBER = "Account Number"
	CAPONE_TRANSACTION_DATE = "Transaction Date"
	CAPONE_TRANSACTION_AMOUNT = "Transaction Amount"
	CAPONE_TRANSACTION_TYPE = "Transaction Type"
	CAPONE_TRANSACTION_DESCRIPTION = "Transaction Description"
	CAPONE_BALANCE = "Balance"
	CAPONE_CREDIT = "Credit"
	CAPONE_DEBIT = "Debit"
	AMEX_DATE = "Date"
	AMEX_DESCRIPTION = "Description"
	AMEX_CARD_MEMBER = "Card Member"
	AMEX_ACCOUNT_NUMBER = "Account #"
	AMEX_AMOUNT = "Amount"
	QUOTE = "\""
	COMMA = ","
)

// this function is 2023 order
// func GetCapitalOneCheckingCSVColumns() []string {
// 	return []string{
// 		CAPONE_ACCOUNT_NUMBER,
// 		CAPONE_TRANSACTION_DATE,
// 		CAPONE_TRANSACTION_AMOUNT,
// 		CAPONE_TRANSACTION_TYPE,
// 		CAPONE_TRANSACTION_DESCRIPTION,
// 		CAPONE_BALANCE,
// 	}
// }

// 2024 order ->
// Account Number,Transaction Description,Transaction Date,Transaction Type,Transaction Amount,Balance
func GetCapitalOneCheckingCSVColumns() []string {
	return []string{
		CAPONE_ACCOUNT_NUMBER,
		CAPONE_TRANSACTION_DESCRIPTION,
		CAPONE_TRANSACTION_DATE,
		CAPONE_TRANSACTION_TYPE,
		CAPONE_TRANSACTION_AMOUNT,
		CAPONE_BALANCE,
	}
}

type CapOneTransaction struct {
	AccountNumber string
	TransactionDate string
	TransactionAmount string
	TransactionType string
	TransactionDescription string
	Balance string
}

// NOTE: ASSUMES ASCII! some characters in full utf-8 may break
func SplitOnComma(line string) ([]string, error) {

	hasQuotes := strings.Contains(line, QUOTE)

	if !hasQuotes {
		return strings.Split(line, ","), nil
	}

	firstQuoteIndex := strings.Index(line, QUOTE)
	secondQuoteIndex := strings.Index(line[firstQuoteIndex + 1:], QUOTE)

	if secondQuoteIndex == 0 {
		return []string{}, errors.New("no second quote in line!")
	}

	columnCount := 0
	lastCommaIndex := 0
	columns := []string{}
	var insideQuotes bool

	for index, rune := range line {
		char := string(rune)
		//log.Println("char:", char)
		if char == QUOTE {
			insideQuotes = !insideQuotes
		}
		if char == COMMA && !insideQuotes {
			column := line[lastCommaIndex:index]
			//log.Println("column:", column)
			columns = append(columns, column)
			columnCount += 1
			lastCommaIndex = index + 1
			continue
		}
		// ASCII
	}
	column := line[lastCommaIndex:]
	columns = append(columns, column)	

	lines := []string{}
	for _, columnEntries := range columns{
		lines = append(lines, columnEntries)
	}
	return lines, nil
}

func (t CapOneTransaction) ToExpenditure(userId *int) (*ncsql.Expenditure, error) {
	e := ncsql.Expenditure{}

	value, err := strconv.ParseFloat(t.TransactionAmount, 32)

	if err != nil {
		log.Println("CapOneTransaction.ToExpenditure() failed to parse value:", t.TransactionAmount)
		return nil, err
	}

	asFloat32 := float32(value)

	isCredit := t.TransactionType == CAPONE_CREDIT

	if isCredit && asFloat32 > 0 {
		asFloat32 *= -1
	}

	description := t.TransactionDescription

	e.Value = &asFloat32
	e.UserId = userId
	e.Description = &description
	e.DateOccurred = &t.TransactionDate

	return &e, nil
}

func StripEmptyLines(lines []string) []string {
	strippedLines := []string{}

	for _, line := range lines {
		if line != "" {
			strippedLines = append(strippedLines, line)
		}
	}
	return strippedLines
}

func ParseCapitalOneCSV(fileBytes []byte) ([]CapOneTransaction, error) {

	// fileBytes, readError := os.ReadFile(path)

	// if readError != nil {
	// 	return nil, readError
	// }

	fileAsString := string(fileBytes)

	fileLines := strings.Split(fileAsString, "\n")

	if len(fileLines) < 2 {
		return []CapOneTransaction{}, nil
	}

	dataLines := StripEmptyLines(fileLines[1:])

	if len(dataLines) == 0 {
		return []CapOneTransaction{}, nil
	}

	transactionCount := len(dataLines)

	transactions := make([]CapOneTransaction, transactionCount, transactionCount)

	for index, line := range dataLines {

		transaction := CapOneTransaction{}

		columnsOnLine, err := SplitOnComma(line)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		for columnIndex, column := range columnsOnLine {

			columnNames := GetCapitalOneCheckingCSVColumns()

			columnName := columnNames[columnIndex]
			// transaction[columnName] = column

			if columnName == CAPONE_ACCOUNT_NUMBER {
				transaction.AccountNumber = column
			} else if columnName == CAPONE_TRANSACTION_DATE {
				transaction.TransactionDate = column
			} else if columnName == CAPONE_TRANSACTION_AMOUNT {
				transaction.TransactionAmount = column
			} else if columnName == CAPONE_TRANSACTION_TYPE {
				transaction.TransactionType = column
			} else if columnName == CAPONE_TRANSACTION_DESCRIPTION {
				transaction.TransactionDescription = column
			} else if columnName == CAPONE_BALANCE {
				transaction.Balance = column
			}
			// transactions[index] = transaction
		}
		transactions[index] = transaction
	}
	return transactions, nil
}

type CustomAmexCheckingTransaction struct {
	Date string
	Description string
	CardMember string
	AccountNumber string
	Amount string
}

func GetAmexCardCSVColumns() []string {
	return []string{
		AMEX_DATE,
		AMEX_DESCRIPTION,
		AMEX_CARD_MEMBER,
		AMEX_ACCOUNT_NUMBER,
		AMEX_AMOUNT,
	}
}

type AmexTransaction struct {
	Date string
	Description string
	CardMember string
	AccountNumber string
	Amount string
}

func (t AmexTransaction) ToExpenditure(userId *int) (*ncsql.Expenditure, error) {
	e := ncsql.Expenditure{}

	value, err := strconv.ParseFloat(t.Amount, 32)

	if err != nil {
		log.Println("AmexTransaction.ToExpenditure() failed to parse value:", t.Amount)
		return nil, err
	}

	asFloat32 := float32(value)
	description := t.Description

	e.Value = &asFloat32
	e.UserId = userId
	e.Description = &description
	e.DateOccurred = &t.Date

	return &e, nil
}

func ParseAmexCreditCardCSV(fileBytes []byte) ([]AmexTransaction, error) {

	// read file
	// fileBytes, readError := os.ReadFile(path)

	// if readError != nil {
	// 	return nil, readError
	// }

	fileAsString := string(fileBytes)

	log.Println("fileAsString:", fileAsString)

	fileLines := strings.Split(fileAsString, "\n")

	if len(fileLines) < 2 {
		return []AmexTransaction{}, nil
	}

	dataLines := StripEmptyLines(fileLines[1:])

	if len(dataLines) == 0 {
		return []AmexTransaction{}, nil
	}

	transactionCount := len(dataLines)

	transactions := make([]AmexTransaction, transactionCount, transactionCount)

	for index, line := range dataLines {

		transaction := AmexTransaction{}

		columns, err := SplitOnComma(line)

		if err != nil {
			log.Println(err.Error(), "skipping file index", index, "in amex csv")
			continue
		}

		for columnIndex, column := range columns {

			columnNames := GetAmexCardCSVColumns()

			if len(columnNames) != len(columns) {
				log.Println("column vs column name mis-match at index", index)
				log.Println("columnNames:", columnNames)
				log.Println("len(columnNames)", len(columnNames))
				log.Println("columns:", columns)
				log.Println("len(columns):", len(columns))
			}

			columnName := columnNames[columnIndex]

			if columnName == AMEX_ACCOUNT_NUMBER {
				transaction.AccountNumber = column
			} else if columnName == AMEX_DATE {
				transaction.Date = column
			} else if columnName == AMEX_AMOUNT {
				transaction.Amount = column
			} else if columnName == AMEX_DESCRIPTION {
				transaction.Description = column
			}
			//transactions[index] = transaction
		}
		transactions[index] = transaction
	}
	return transactions, nil
}

func CapOneTransactionsToExpenditures(transactions []CapOneTransaction, userId *int) ([]ncsql.Expenditure, error) {
	size := len(transactions)

	expenditures := make([]ncsql.Expenditure, size, size)

	var index int

	for _, transaction := range transactions {
		expenditure, err := transaction.ToExpenditure(userId)
		if err != nil {
			log.Println("error calling transaction.ToExpenditure():", err.Error())
			log.Println("error was at index", index)
			continue
		}
		log.Println("assigning index", index)
		expenditures[index] = *expenditure
		index++
	}

	return expenditures, nil
}

func AmexTransactionsToExpenditures(transactions []AmexTransaction, userId *int) ([]ncsql.Expenditure, error) {
	size := len(transactions)

	log.Println("size", size)

	expenditures := make([]ncsql.Expenditure, size, size)

	var index int

	for _, transaction := range transactions {
		expenditure, err := transaction.ToExpenditure(userId)
		if err != nil {
			log.Println("error calling transaction.ToExpenditure():", err.Error())
			log.Println("error was at index", index)
			continue
		}
		log.Println("assigning index", index)
		expenditures[index] = *expenditure
		index++
	}

	return expenditures, nil
}
