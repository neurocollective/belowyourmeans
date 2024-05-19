package parsing

import (
	"os"
	//"fmt"
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
)

func GetCapitalOneCheckingCSVColumns() []string {
	return []string{
		CAPONE_ACCOUNT_NUMBER,
		CAPONE_TRANSACTION_DATE,
		CAPONE_TRANSACTION_AMOUNT,
		CAPONE_TRANSACTION_TYPE,
		CAPONE_TRANSACTION_DESCRIPTION,
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

func (t CapOneTransaction) ToExpenditure(userId *int) (*ncsql.Expenditure, error) {
	e := ncsql.Expenditure{}

	value, err := strconv.ParseFloat(t.TransactionAmount, 32)

	if err != nil {
		log.Println("CapOneTransactionsToExpenditures failed to parse value:", t.TransactionAmount)
		return nil, err
	}

	asFloat32 := float32(value)
	description := t.TransactionType + ": " + t.TransactionDescription

	e.Value = &asFloat32
	e.UserId = userId
	e.Description = &description
	e.DateOccurred = &t.TransactionDate

	return &e, nil
}

func ParseCapitalOneCSV(path string) ([]CapOneTransaction, error) {

	fileBytes, readError := os.ReadFile(path)

	if readError != nil {
		return nil, readError
	}

	fileAsString := string(fileBytes)

	fileLines := strings.Split(fileAsString, "\n")

	transactionCount := len(fileLines)

	transactions := make([]CapOneTransaction, transactionCount, transactionCount)

	for index, line := range fileLines {

		transaction := CapOneTransaction{}

		for columnIndex, column := range strings.Split(line, ",") {

			columnsNames := GetCapitalOneCheckingCSVColumns()

			columnName := columnsNames[columnIndex]
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
			transactions[index] = transaction
		}
	}
	return transactions[1:], nil
}

func GetCustomAmexCheckingCSVColumns() []string {
	return []string{
		"Date",
		"Description",
		"Credits",
		"Debits",
		"Balance",
	}
}

type CustomAmexCheckingTransaction struct {
	Date string
	Description string
	CardMember string
	AccountNumber string
	Amount string
}

func ParseCustomAmexCheckingCSV(path string) ([]AmexTransaction, error) {

	// read file
	// fileBytes, readError := os.ReadFile(path)
	// split on newline
	// split on commas

	return []AmexTransaction{}, nil
}

func GetAmexCardCSVColumns() []string {
	return []string{
		"Date",
		"Description",
		"Card Member",
		"Account #",
		"Amount",
	}
}

type AmexTransaction struct {
	Date string
	Description string
	CardMember string
	AccountNumber string
	Amount string
}

func ParseAmexCreditCardCSV(path string) ([]AmexTransaction, error) {

	// read file
	// fileBytes, readError := os.ReadFile(path)
	// split on newline
	// split on commas

	return []AmexTransaction{}, nil
}

func CapOneTransactionsToExpenditures(transactions []CapOneTransaction, userId *int) ([]ncsql.Expenditure, error) {
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