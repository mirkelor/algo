package blind75

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if left > right || top > bottom {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
