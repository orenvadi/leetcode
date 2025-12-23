package main

import (
	"fmt"
)

var tests = []struct {
	grid       [][]int
	minPathSum int
}{
	{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
	{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
}

func main() {
	for _, testCase := range tests {
		res := uniquepathii.minPathSum(testCase.grid)
		if res != testCase.minPathSum {
			fmt.Printf("minPathSum(%v) = %d, expected %d", testCase.grid, res, testCase.minPathSum)
		}
	}
}
