package utils

import "testing"

func TestReadFile(t *testing.T) {
	texts:=ReadFile("../data/comments100.txt")
	if len(texts)!=100{
		t.Error("Error line amount read from file")
	}
}