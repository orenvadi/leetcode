package main

import "testing"

var tests = []struct {
	testCase any
	exp      any
}{}

func Test(t *testing.T) {
	for _, testCase := range tests {
		res := some(testCase.testCase)
		if res != testCase.exp {
			t.Errorf("expWithObstacles(%v) = %v, expected %t", testCase.testCase, res, testCase.exp)
		}
	}
}
