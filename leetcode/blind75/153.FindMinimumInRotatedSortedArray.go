package blind75

import "github.com/Mirkelor/algo/common"

func findMin(nums []int) int {
	l, r, res := 0, len(nums)-1, nums[0]
	for l <= r {
		if nums[l] <= nums[r] {
			res = common.Min(res, nums[l])
			break
		}
		mid := l + (r-l)/2
		res = common.Min(res, nums[mid])
		if nums[mid] >= nums[l] {
			l = mid + 1
			res = common.Min(res, nums[l])
		} else {
			r = mid - 1
		}
	}
	return res
}
