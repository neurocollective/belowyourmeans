package main

import (
	"log"
	// "neurocollective.io/neurocollective/belowyourmeans/src/db"
	// "neurocollective.io/neurocollective/belowyourmeans/src/constants"
	"neurocollective.io/neurocollective/belowyourmeans/src/parsing"
	"os"
	ncsql "github.com/neurocollective/go_utils/sql"
)

func main() {

	client, getClientError := ncsql.BuildPostgresClient("user=postgres password=postgres dbname=postgres sslmode=disable")		

	// connect to the db to test if connection is valid

	if getClientError != nil {
		log.Fatal("error getting client")
	}

	cwd, _ := os.Getwd()

	log.Println("cwd", cwd)

	// transactions == []parsing.CapOneTransaction
	transactions, parseError := parsing.ParseCapitalOneCSV(cwd + "/sample_files/capone_checking_2023_11_5.csv")

	if parseError != nil {
		log.Fatal("error!" + parseError.Error())
	}

	if len(transactions) > 10 {
		log.Println("first ten transactions:")
		log.Println(transactions[:9])
	} else {
		log.Println(transactions)		
	}

	tCount := len(transactions)

	log.Println("transactions in total:", tCount)

	insertQueryStem := "INSERT INTO expenditure VALUES ("

	valuesCount := len(transactions) * 6 // size of transactions, times 6 fields in each transaction

	args := make([]any, valuesCount, valuesCount)

	for _ , capOneTransaction := range transactions {


	sql.SimpleQuery(client, query, args)

		// type CapOneTransaction struct {
		// 	AccountNumber string
		// 	TransactionDate string
		// 	TransactionAmount string
		// 	TransactionType string
		// 	TransactionDescription string
		// 	Balance string
		// }

	}

}