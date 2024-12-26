package parsing

import (
	"testing"
	"os"
)

func TestSplitOnCommaOne(t *testing.T) {

	line := "0452,\"ATM Withdrawal - WALGREENS # -X00 AX002278 BROOKLYN, NY\",12/19/24,Debit,300,10338.61"
	columns, err := SplitOnComma(line)

	if err != nil {
		t.Fatal("expected no err, got an err: " + err.Error())
	}

	expectedOne := "0452"
	expectedTwo := "\"ATM Withdrawal - WALGREENS # -X00 AX002278 BROOKLYN, NY\""
	expectedThree := "12/19/24"
	expectedFour := "Debit"
	expectedFive := "300"
	expectedSix := "10338.61"

	if len(columns) < 6 {
		t.Fatal("too few columns")		
	}

	if columns[0] != expectedOne {
		t.Fatalf("expected first column to be %v but got %v", expectedOne, columns[0])
	}
	if columns[1] != expectedTwo {
		t.Fatalf("expected second column to be %v but got %v", expectedTwo, columns[1])
	}
	if columns[2] != expectedThree {
		t.Fatalf("expected third column to be %v but got %v", expectedThree, columns[2])
	}
	if columns[3] != expectedFour {
		t.Fatalf("expected fourth column to be %v but got %v", expectedFour, columns[3])
	}
	if columns[4] != expectedFive {
		t.Fatalf("expected fifth column to be %v but got %v", expectedFive, columns[4])
	}
	if columns[5] != expectedSix {
		t.Fatalf("expected sixth column to be %v but got %v", expectedSix, columns[5])
	}
}

func TestSplitOnCommaTwo(t *testing.T) {

	line := "0452,Withdrawal from VENMO PAYMENT,12/20/24,Debit,1500,8838.61"
	columns, err := SplitOnComma(line)

	if err != nil {
		t.Fatal("expected no err, got an err: " + err.Error())
	}

	expectedOne := "0452"
	expectedTwo := "Withdrawal from VENMO PAYMENT"
	expectedThree := "12/20/24"
	expectedFour := "Debit"
	expectedFive := "1500"
	expectedSix := "8838.61"

	if len(columns) < 6 {
		t.Fatal("too few columns")		
	}

	if columns[0] != expectedOne {
		t.Fatalf("expected first column to be %v but got %v", expectedOne, columns[0])
	}
	if columns[1] != expectedTwo {
		t.Fatalf("expected second column to be %v but got %v", expectedTwo, columns[1])
	}
	if columns[2] != expectedThree {
		t.Fatalf("expected third column to be %v but got %v", expectedThree, columns[2])
	}
	if columns[3] != expectedFour {
		t.Fatalf("expected fourth column to be %v but got %v", expectedFour, columns[3])
	}
	if columns[4] != expectedFive {
		t.Fatalf("expected fifth column to be %v but got %v", expectedFive, columns[4])
	}
	if columns[5] != expectedSix {
		t.Fatalf("expected sixth column to be %v but got %v", expectedSix, columns[5])
	}
}

func TestSplitOnCommaThree(t *testing.T) {

	line := "1,\"Dude, sup\",3,4,5,6"
	columns, err := SplitOnComma(line)

	if err != nil {
		t.Fatal("expected no err, got an err: " + err.Error())
	}

	expectedOne := "1"
	expectedTwo := "\"Dude, sup\""
	expectedThree := "3"
	expectedFour := "4"
	expectedFive := "5"
	expectedSix := "6"

	if len(columns) < 6 {
		t.Fatal("too few columns")		
	}

	if columns[0] != expectedOne {
		t.Fatalf("expected first column to be %v but got %v", expectedOne, columns[0])
	}
	if columns[1] != expectedTwo {
		t.Fatalf("expected second column to be %v but got %v", expectedTwo, columns[1])
	}
	if columns[2] != expectedThree {
		t.Fatalf("expected third column to be %v but got %v", expectedThree, columns[2])
	}
	if columns[3] != expectedFour {
		t.Fatalf("expected fourth column to be %v but got %v", expectedFour, columns[3])
	}
	if columns[4] != expectedFive {
		t.Fatalf("expected fifth column to be %v but got %v", expectedFive, columns[4])
	}
	if columns[5] != expectedSix {
		t.Fatalf("expected sixth column to be %v but got %v", expectedSix, columns[5])
	}
}

func TestParseCapitalOneCSV(t *testing.T) {

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err.Error())
	}

	transactions, err := ParseCapitalOneCSV(cwd + "/capone_test.csv")

	if err != nil {
		t.Fatal("got an error running ParseCapitalOneCSV: " + err.Error())
	}

	if len(transactions) != 1 {
		t.Fatalf("expected transactions to be length 1 but got %d", len(transactions))	
	}

	first := transactions[0]

	if first.AccountNumber != "0452" {
		t.Fatalf("expected transactions to be `0452` but got %v", first.AccountNumber)	
	}

	if first.TransactionDate != "12/20/24" {
		t.Fatalf("expected transactions to be `12/20/24` but got %v", first.TransactionDate)	
	}

	if first.TransactionAmount != "11.43" {
		t.Fatalf("expected transactions to be `11.43` but got %v", first.TransactionAmount)	
	}

	if first.TransactionType != "Debit" {
		t.Fatalf("expected transactions to be `Debit` but got %v", first.TransactionType)	
	}

	if first.TransactionDescription != "Debit Card Purchase - SQ GS COFFEE BROOKLYN NY" {
		t.Fatalf("expected transactions to be `Debit Card Purchase - SQ GS COFFEE BROOKLYN NY` but got %v", first.TransactionDescription)	
	}

	if first.Balance != "8827.18" {
		t.Fatalf("expected transactions to be `8827.18` but got %v", len(transactions))	
	}

}
