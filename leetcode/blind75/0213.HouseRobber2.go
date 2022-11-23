package blind75

import "github.com/Mirkelor/algo/common"

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return common.Max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	profit, prev1, prev2 := 0, 0, 0
	for _, cur := range nums {
		profit = common.Max(prev2+cur, prev1)
		prev2, prev1 = prev1, profit
	}
	return profit
}
