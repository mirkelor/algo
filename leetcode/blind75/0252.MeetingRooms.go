package blind75

import (
	"sort"

	"github.com/Mirkelor/algo/common"
)

func canAttendMeetings(intervals []*common.Interval) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1].End > intervals[i].Start {
			return false
		}
	}
	return true
}
