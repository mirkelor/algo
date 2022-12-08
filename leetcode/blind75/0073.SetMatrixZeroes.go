package blind75

func setZeroes(matrix [][]int) {
	isFirstRowZero := false
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					isFirstRowZero = true
				}
			}
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for r := range matrix {
			matrix[r][0] = 0
		}
	}

	if isFirstRowZero {
		for c := range matrix[0] {
			matrix[0][c] = 0
		}
	}
}
