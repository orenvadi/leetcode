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
		res := plusOne(testCase.testCase)
		if slices.Equal(res, testCase.exp) {
			t.Errorf("expWithObstacles(%v) = %v, expected %v", testCase.testCase, res, testCase.exp)
		}
	}
}
