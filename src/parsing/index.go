package parsing

import (
	"os"
	"fmt"
	"strings"
)

func GetCapitalOneCheckingCSVColumns() []string {
	return []string{
		"Account Number",
		"Transaction Date",
		"Transaction Amount",
		"Transaction Type",
		"Transaction Description",
		"Balance",
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

func ParseCapitalOneCSV(path string) ([]CapOneTransaction, error) {

	fileBytes, readError := os.ReadFile(path)

	if readError != nil {
		return nil, readError
	}

	fileAsString := string(fileBytes)

	fileLines := strings.Split(fileAsString, "\n")

	for lineIndex, line := range fileLines {

		transaction := make(map[string]string)

		for columnIndex, column := range strings.Split(line, ",") {

			columnsNames := GetCapitalOneCheckingCSVColumns()

			columnName := columnsNames[columnIndex]
			transaction[columnName] = column
		}

		// just first line for now
		if lineIndex > 0 {
			fmt.Println("transaction", transaction)
			break
		}
	}

	// split on newline
	// split on commas
	return []CapOneTransaction{}, nil
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

func ParseAmericanExpressCreditCardCSV(path string) ([]AmexTransaction, error) {

	// read file
	// fileBytes, readError := os.ReadFile(path)
	// split on newline
	// split on commas

	return []AmexTransaction{}, nil
}