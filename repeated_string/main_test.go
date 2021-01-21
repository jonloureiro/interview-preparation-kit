package main

import "testing"

func TestRepeatedString(t *testing.T) {
	test := repeatedString("aba", 10)
	var result int64
	result = 7
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = repeatedString("a", 1000000000000)
	result = 1000000000000
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
