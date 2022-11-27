package blind75

func numDecodings(s string) int {
	n := len(s)
	dp := map[int]int{n: 1}

	var dfs func(int) int
	dfs = func(i int) int {
		if d, ok := dp[i]; ok {
			return d
		}
		if s[i] == '0' {
			return 0
		}

		res := dfs(i + 1)
		if i+1 < n && (s[i] == '1' || s[i] == '2' && s[i+1] <= '6') {
			res += dfs(i + 2)
		}
		dp[i] = res
		return res
	}

	return dfs(0)
}
