package blind75

import (
	"sort"

	"github.com/Mirkelor/algo/common"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res, prevEnd := 0, intervals[0][1]
	for _, cur := range intervals[1:] {
		if cur[0] >= prevEnd {
			prevEnd = cur[1]
		} else {
			res++
			prevEnd = common.Min(prevEnd, cur[1])
		}
	}

	return res
}
