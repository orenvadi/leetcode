package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	testCase []int
	exp      []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	{[]int{9}, []int{1, 0}},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		var tCase []int
		copy(tCase, testCase.testCase)

		res := plusOne(tCase)
		if slices.Equal(res, testCase.exp) {
			t.Errorf("expWithObstacles(%v) = %v, expected %v", tCase, res, testCase.exp)
		}
	}
}
