package blind75

import "github.com/mirkelor/algo/common"

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for i, arr := range intervals {
		if newInterval[1] < arr[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > arr[1] {
			res = append(res, arr)
		} else {
			newInterval = []int{common.Min(newInterval[0], arr[0]), common.Max(newInterval[1], arr[1])}
		}
	}

	res = append(res, newInterval)
	return res
}
