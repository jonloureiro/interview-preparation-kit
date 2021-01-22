package main

import "testing"

func TestHourglassSum(t *testing.T) {
	test := hourglassSum([][]int32{
		{
			1, 1, 1, 0, 0, 0,
		},
		{
			0, 1, 0, 0, 0, 0,
		},
		{
			1, 1, 1, 0, 0, 0,
		},
		{
			0, 0, 2, 4, 4, 0,
		},
		{
			0, 0, 0, 2, 0, 0,
		},
		{
			0, 0, 1, 2, 4, 0,
		},
	})
	var result int32
	result = 19

	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
