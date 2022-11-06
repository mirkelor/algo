package blind75

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {
		res[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(nums) - 1; i > 0; i-- {
		product *= nums[i]
		res[i-1] *= product
	}

	return res
}
