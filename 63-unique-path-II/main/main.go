package main

import (
	"fmt"

	uniquepathii "github.com/orenvadi/leetcode/63-unique-path-II"
)

var tests = []struct {
	obstacleGrid [][]int
	uniquePaths  int
}{
	{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
	{[][]int{{0, 0}, {1, 1}, {0, 0}}, 0},
	{[][]int{{0, 1}, {0, 0}}, 1},
	{[][]int{{0, 0}, {0, 1}}, 0},
	{[][]int{{0, 0}, {1, 0}}, 1},
	{
		[][]int{
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{1, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{1, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 1},
			{0, 0},
			{0, 0},
			{1, 0},
			{0, 0},
			{0, 0},
			{0, 1},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 1},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{1, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
		}, 0,
	},
}

func main() {
	for _, testCase := range tests {
		res := uniquepathii.UniquePathsWithObstacles(testCase.obstacleGrid)
		// if res != testCase.uniquePaths {
		fmt.Printf("uniquePathsWithObstacles(%v) = %d, expected %d\n", testCase.obstacleGrid, res, testCase.uniquePaths)
		// }
	}
}
