package blind75

func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1
	for l < r {
		for i := 0; i < r-l; i++ {
			matrix[l][l+i], matrix[l+i][r], matrix[r][r-i], matrix[r-i][l] = matrix[r-i][l], matrix[l][l+i], matrix[l+i][r], matrix[r][r-i]
		}
		l++
		r--
	}
}
