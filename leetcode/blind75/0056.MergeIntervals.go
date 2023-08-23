package blind75

import (
	"sort"

	"github.com/mirkelor/algo/common"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res := make([][]int, 0)
	res = append(res, intervals[0])
	for _, interval := range intervals {
		if res[len(res)-1][1] >= interval[0] {
			res[len(res)-1][1] = common.Max(res[len(res)-1][1], interval[1])
		} else {
			res = append(res, interval)
		}
	}

	return res
}
