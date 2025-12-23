package main

import "testing"

var tests = []struct {
	testCase string
	exp      bool
}{}

func Test(t *testing.T) {
	for _, testCase := range tests {
		res := isNumber(testCase.testCase)
		if res != testCase.exp {
			t.Errorf("expWithObstacles(%v) = %v, expected %t", testCase.testCase, res, testCase.exp)
		}
	}
}
