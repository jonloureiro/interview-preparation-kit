package main

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	test := countingValleys(8, "UDDDUDUU")
	var result int32
	result = 1
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = countingValleys(12, "DDUUDDUDUUUD")
	result = 2
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
