package blind75

import "github.com/mirkelor/algo/common"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	longest := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = common.Max(dp[i], 1+dp[j])
			}
		}
		longest = common.Max(longest, dp[i])
	}
	return longest
}
