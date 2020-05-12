package leetcode

func MinimumPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	for i := range grid {
		for j := range grid[i] {
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
				continue
			}
			if i > 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if j > 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
	// Brute Force
	// res := 1<<63 - 1
	// MinimimPathSumDfs(grid, &res, 0, 0, 0)
	// return res

}

func min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}

// Brute-Force algorithm
func MinimimPathSumDfs(grid [][]int, res *int, sum, i, j int) {
	sum += grid[i][j]
	if i == len(grid)-1 && j == len(grid[0])-1 {
		if *res > sum {
			*res = sum
		}
	}
	if i+1 < len(grid) {
		MinimimPathSumDfs(grid, res, sum, i+1, j)
	}
	if j+1 < len(grid[0]) {
		MinimimPathSumDfs(grid, res, sum, i, j+1)
	}
	sum -= grid[i][j]
}
