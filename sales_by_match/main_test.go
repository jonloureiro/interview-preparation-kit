package main

import "testing"

func TestSockMerchant(t *testing.T) {
	test := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	var result int32
	result = 3
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2})
	result = 2
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = sockMerchant(1, []int32{1})
	result = 0
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = sockMerchant(3, []int32{1, 2, 3})
	result = 0
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
