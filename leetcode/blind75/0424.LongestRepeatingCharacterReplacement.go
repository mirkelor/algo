package blind75

import "github.com/mirkelor/algo/common"

func characterReplacement(s string, k int) int {
	l, res, maxf, count := 0, 0, 0, make(map[byte]int)
	for r := range s {
		count[s[r]]++
		maxf = common.Max(maxf, count[s[r]])
		for r-l+1-maxf > k {
			count[s[l]]--
			l++
		}
		res = common.Max(res, r-l+1)
	}
	return res
}
