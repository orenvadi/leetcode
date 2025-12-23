package uniquepathii

func minPathSum(grid [][]int) int {
	gridLenghtM := len(grid)
	gridLenghtN := len(grid[0])

	for i, row := range grid {
		for j, num := range row {
			if i == 0 && j == 0 {
				continue
			}

			left := 999999
			if i > 0 {
				left = grid[i-1][j]
			}

			right := 999999
			if j > 0 {
				right = grid[i][j-1]
			}

			grid[i][j] = num + min(left, right)
		}
	}

	result := grid[gridLenghtM-1][gridLenghtN-1]

	return result
}
