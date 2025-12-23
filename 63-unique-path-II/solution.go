package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	obstacleGridLenghtM := len(obstacleGrid)
	obstacleGridLenghtN := len(obstacleGrid[0])

	// first is obstacle
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// if last is obstacle
	if obstacleGrid[obstacleGridLenghtM-1][obstacleGridLenghtN-1] == 1 {
		return 0
	}

	for i, row := range obstacleGrid {
		for j, num := range row {
			switch {

			case i == 0 && j == 0:
				obstacleGrid[0][0] = 1

			case num == 1:
				obstacleGrid[i][j] = 0

			default:
				left := 0
				if i > 0 {
					left = obstacleGrid[i-1][j]
				}
				right := 0
				if j > 0 {
					right = obstacleGrid[i][j-1]
				}
				obstacleGrid[i][j] = left + right
			}
		}
	}

	result := obstacleGrid[obstacleGridLenghtM-1][obstacleGridLenghtN-1]

	return result
}
