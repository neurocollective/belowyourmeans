package parsing

import (
	"testing"
	// "os"
	// "log"
	// ncsql "github.com/neurocollective/go_utils/sql"
)

// func TestSplitOnCommaOne(t *testing.T) {

// 	line := "0452,\"ATM Withdrawal - WALGREENS # -X00 AX002278 BROOKLYN, NY\",12/19/24,Debit,300,10338.61"
// 	columns, err := SplitOnComma(line)

// 	if err != nil {
// 		t.Fatal("expected no err, got an err: " + err.Error())
// 	}

// 	if columns[0] != "0452" {
// 		t.Fatalf("expected first column to be %v but got %v", "0452", columns[0])
// 	}
// }

// func TestSplitOnCommaTwo(t *testing.T) {

// 	line := "0452,Withdrawal from VENMO PAYMENT,12/20/24,Debit,1500,8838.61"
// 	columns, err := SplitOnComma(line)

// 	if err != nil {
// 		t.Fatal("expected no err, got an err: " + err.Error())
// 	}

// 	if columns[0] != "0452" {
// 		t.Fatalf("expected first column to be %v but got %v", "0452", columns[0])
// 	}
// }

func TestSplitOnCommaThree(t *testing.T) {

	line := "1,\"Dude, sup\",3,4,5,6"
	columns, err := SplitOnComma(line)

	if err != nil {
		t.Fatal("expected no err, got an err: " + err.Error())
	}

	if columns[0] != "1" {
		t.Fatalf("expected first column to be %v but got %v", "1", columns[0])
	}
}