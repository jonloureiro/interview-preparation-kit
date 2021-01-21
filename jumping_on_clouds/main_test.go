package main

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	test := jumpingOnClouds([]int32{0, 0, 1, 0, 1, 0, 0})
	var result int32
	result = 4
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}

	test = jumpingOnClouds([]int32{0, 1, 0, 0, 0, 1, 0})
	result = 3
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
