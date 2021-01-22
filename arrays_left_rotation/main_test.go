package main

import (
	"reflect"
	"testing"
)

func TestRotLeft(t *testing.T) {
	test := rotLeft([]int32{1, 2, 3, 4, 5}, 4)
	result := []int32{5, 1, 2, 3, 4}

	if !reflect.DeepEqual(test, result) {
		t.Error("Expected:", result, "Got:", test)
	}
}
