package blind75

import "github.com/mirkelor/algo/common"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for a := 1; a < amount+1; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				dp[a] = common.Min(dp[a], 1+dp[a-c])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
