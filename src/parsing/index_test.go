package parsing

import (
	"testing"
	"os"
	"log"
	ncsql "github.com/neurocollective/go_utils/sql"
)

func TestParseCapitalOneCSV(t *testing.T) {

	client, getClientError := ncsql.BuildPostgresClient("user=postgres password=postgres dbname=postgres sslmode=disable")		

	// connect to the db to test if connection is valid

	if getClientError != nil {
		log.Fatal("error getting client")
	}

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