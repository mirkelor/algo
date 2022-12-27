package blind75

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if len(word) > m*n {
		return false
	}

	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var backtrack func(r, c, i int) bool
	backtrack = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		rowInbound := 0 <= r && r < m
		colInbound := 0 <= c && c < n
		if !rowInbound || !colInbound || visit[r][c] || board[r][c] != word[i] {
			return false
		}

		visit[r][c] = true
		res := backtrack(r+1, c, i+1) || backtrack(r-1, c, i+1) || backtrack(r, c+1, i+1) || backtrack(r, c-1, i+1)
		visit[r][c] = false
		return res
	}

	for r := range board {
		for c := range board[r] {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false
}
