package blind75

import (
	"sort"

	"github.com/Mirkelor/algo/common"
)

func minMeetingRooms(intervals []*common.Interval) int {
	start, end := intervals, intervals
	sort.Slice(start, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	sort.Slice(end, func(i, j int) bool { return intervals[i].End < intervals[j].End })

	res, count := 0, 0
	for len(start) > 0 && len(end) > 0 {
		if start[0].Start < end[0].End {
			count++
			start = start[1:]
		} else {
			count--
			end = end[1:]
		}
		res = common.Max(res, count)
	}

	return res
}
