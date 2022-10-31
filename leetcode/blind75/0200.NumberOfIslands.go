package blind75

func numIslands(grid [][]byte) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				dfsNumIslands(r, c, grid)
				count++
			}
		}
	}
	return count
}

func dfsNumIslands(r int, c int, grid [][]byte) {
	rowInbound := 0 <= r && r < len(grid)
	colInbound := 0 <= c && c < len(grid[0])
	if !rowInbound || !colInbound || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	dfsNumIslands(r+1, c, grid)
	dfsNumIslands(r-1, c, grid)
	dfsNumIslands(r, c+1, grid)
	dfsNumIslands(r, c-1, grid)
}
