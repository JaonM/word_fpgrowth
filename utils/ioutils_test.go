package utils

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	texts := ReadFile("../test/comments100.txt")
	if len(texts) != 100 {
		t.Error("Error line amount read from file")
	}
}

func TestWriteCSV(t *testing.T) {
	WriteCSV("../test/test.csv", [] string{"name", "gender"}, [][]string{{"Alice", "female"}, {"Bob", "male"}})
	WriteCSV("../test/test2.csv", nil, [][]string{{"Alice", "female"}, {"Bob", "male"}})
}
