package blind75

import "github.com/Mirkelor/algo/common"

func maxProduct(nums []int) int {
	l, r, res := 0, 0, nums[0]
	for i := range nums {
		if l == 0 {
			l = 1
		}
		if r == 0 {
			r = 1
		}
		l *= nums[i]
		r *= nums[len(nums)-1-i]
		res = common.Max(res, common.Max(l, r))
	}
	return res
}
