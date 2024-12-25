package main

import (
	ncsql "github.com/neurocollective/go_utils/sql"
	"log"
	"neurocollective.io/neurocollective/belowyourmeans/server/parsing"
	"os"
)

func main() {

	// args := os.Args

	// fileName := args[1]

	client, getClientError := ncsql.BuildPostgresClient("user=postgres password=postgres dbname=postgres sslmode=disable")

	// connect to the db to test if connection is valid

	if getClientError != nil {
		log.Fatal("error getting client")
	}

	cwd, _ := os.Getwd()

	transactions, err := parsing.ParseCapitalOneCSV(cwd + "/sample_files/cap_one_most_of_2024.csv")

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

	// if len(transactions) > 10 {
	// 	log.Println("trimming transactions")
	// 	transactions = transactions[:9]

	// 	for _, transaction := range transactions {
	// 		log.Println(transaction.TransactionDate)
	// 		log.Println(transaction.TransactionDate)
	// 	}
	// } else {
	// 	log.Println(transactions)
	// }

	userId := 1

	expenditures, err := parsing.CapOneTransactionsToExpenditures(transactions, &userId)

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

	insert := ncsql.Insert[ncsql.Expenditure]

	err = insert(client, expenditures)

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

	amexTransactions, err := parsing.ParseAmexCreditCardCSV(cwd + "/../../sample_files/amex_2024_november.csv")

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

	expenditures, err = parsing.AmexTransactionsToExpenditures(amexTransactions, &userId)

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

	insert = ncsql.Insert[ncsql.Expenditure]

	err = insert(client, expenditures)

	if err != nil {
		log.Fatal("error!" + err.Error())
	}

}
