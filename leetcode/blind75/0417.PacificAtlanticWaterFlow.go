package blind75

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific, atlantic := make([][]bool, m), make([][]bool, m)

	for i := range heights {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	for r := range heights {
		dfsPacificAtlantic(r, 0, heights[r][0], pacific, heights)
		dfsPacificAtlantic(r, n-1, heights[r][n-1], atlantic, heights)
	}

	for c := range heights[0] {
		dfsPacificAtlantic(0, c, heights[0][c], pacific, heights)
		dfsPacificAtlantic(m-1, c, heights[m-1][c], atlantic, heights)
	}

	waterFlow := [][]int{}
	for r := range heights {
		for c := range heights[r] {
			if pacific[r][c] && atlantic[r][c] {
				waterFlow = append(waterFlow, []int{r, c})
			}
		}
	}
	return waterFlow
}

func dfsPacificAtlantic(r, c, prev int, visited [][]bool, heights [][]int) {
	rowInbound := 0 <= r && r < len(heights)
	colInbound := 0 <= c && c < len(heights[0])
	if !rowInbound || !colInbound || visited[r][c] || heights[r][c] < prev {
		return
	}
	visited[r][c] = true
	dfsPacificAtlantic(r+1, c, heights[r][c], visited, heights)
	dfsPacificAtlantic(r-1, c, heights[r][c], visited, heights)
	dfsPacificAtlantic(r, c+1, heights[r][c], visited, heights)
	dfsPacificAtlantic(r, c-1, heights[r][c], visited, heights)
}
