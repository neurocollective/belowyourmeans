package parsing

import (
	"testing"
	"os"
)

func TestParseCapitalOneCSV(t *testing.T) {
	cwd, _ := os.Getwd()

	_, parseError := ParseCapitalOneCSV(cwd + "/../../sample_files/capone_checking_2023_11_5.csv")

	if parseError != nil {
		t.Fatal("error!" + parseError.Error())
	}
}