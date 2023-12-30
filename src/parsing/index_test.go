package parsing

import (
	"testing"
	"os"
	"log"
)

func TestParseCapitalOneCSV(t *testing.T) {
	cwd, _ := os.Getwd()

	transactions, parseError := ParseCapitalOneCSV(cwd + "/../../sample_files/capone_checking_2023_11_5.csv")

	if parseError != nil {
		t.Fatal("error!" + parseError.Error())
	}

	if len(transactions) > 10 {
		log.Println(transactions[:9])
	} else {
		log.Println(transactions)		
	}

	// create query to insert transactions

	// run query

	// get all categories & category labels

	// see which transactions ids match to a category id

	// insert categories to relevant transactions
}