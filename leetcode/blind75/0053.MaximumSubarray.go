package blind75

func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]
	for _, num := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += num
		if sum > max {
			max = sum
		}
	}
	return max
}
