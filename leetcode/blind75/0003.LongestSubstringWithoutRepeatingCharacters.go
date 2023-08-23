package blind75

import "github.com/mirkelor/algo/common"

func lengthOfLongestSubstring(s string) int {
	l, max, count := 0, 0, make(map[byte]int)
	for r := range s {
		for count[s[r]] > 0 {
			count[s[l]]--
			l++
		}
		count[s[r]]++
		max = common.Max(max, r-l+1)
	}
	return max
}
