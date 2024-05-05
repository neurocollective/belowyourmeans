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

	transactions, err := ParseCapitalOneCSV(cwd + "/../../sample_files/capone_checking_2023_11_5.csv")

	if err != nil {
		t.Fatal("error!" + err.Error())
	}

	if len(transactions) > 10 {
		log.Println(transactions[:9])
	} else {
		log.Println(transactions)		
	}

	userId := 1

	expenditures, err := CapOneTransactionsToExpenditures(transactions, &userId)

	if err != nil {
		t.Fatal("error!" + err.Error())
	}

	insert := ncsql.Insert[ncsql.Expenditure]

	err = insert(client, expenditures)

	if err != nil {
		t.Fatal("error!" + err.Error())		
	}

	// create query to insert transactions

	// run query

	// get all categories & category labels

	// see which transactions ids match to a category id

	// insert categories to relevant transactions
}