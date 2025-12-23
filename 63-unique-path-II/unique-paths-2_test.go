package uniquepathii

import "testing"

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

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, testCase := range tests {
		res := uniquePathsWithObstacles(testCase.obstacleGrid)
		if res != testCase.uniquePaths {
			t.Errorf("uniquePathsWithObstacles(%v) = %d, expected %d", testCase.obstacleGrid, res, testCase.uniquePaths)
		}
	}
}
