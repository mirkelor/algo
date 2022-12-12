package blind75

func missingNumber(nums []int) int {
	l := len(nums)
	res := l
	for i := 0; i < l; i++ {
		res ^= (i ^ nums[i])
	}
	return res
}
