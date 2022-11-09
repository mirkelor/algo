package blind75

import "github.com/Mirkelor/algo/common"

func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			max = common.Max(max, height[left]*(right-left))
			left++
		} else {
			max = common.Max(max, height[right]*(right-left))
			right--
		}
	}
	return max
}
