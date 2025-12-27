package main

import "testing"

var tests = []struct {
	testCase [2]string
	exp      string
}{
	{[2]string{"11", "1"}, "100"},
	{[2]string{"1010", "1011"}, "10101"},
	{[2]string{"0", "1"}, "1"},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		res := addBinary(testCase.testCase[0], testCase.testCase[1])
		if res != testCase.exp {
			t.Errorf("expWithObstacles(%v) = %v, expected %t", testCase.testCase, res, testCase.exp)
		}
	}
}
