package leetcode

func NumberOfIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				dfsNumberOfIslands(grid, i, j)
			}
		}
	}
	return res
}

func dfsNumberOfIslands(grid [][]byte, i, j int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == '1' {
		grid[i][j] = 0
		dfsNumberOfIslands(grid, i+1, j)
		dfsNumberOfIslands(grid, i-1, j)
		dfsNumberOfIslands(grid, i, j+1)
		dfsNumberOfIslands(grid, i, j-1)
	}
}
