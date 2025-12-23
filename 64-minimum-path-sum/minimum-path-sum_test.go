package uniquepathii

import "testing"

var tests = []struct {
	grid       [][]int
	minPathSum int
}{
	{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
	{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, testCase := range tests {
		res := minPathSum(testCase.grid)
		if res != testCase.minPathSum {
			t.Errorf("minPathSum(%v) = %d, expected %d", testCase.grid, res, testCase.minPathSum)
		}
	}
}
