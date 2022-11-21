package blind75

func climbStairs(n int) int {
	return climbStairsHelper(n, make([]int, n+1))
}

func climbStairsHelper(n int, memo []int) int {
	if n <= 3 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	return memo[n]
}
