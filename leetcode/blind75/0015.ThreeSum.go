package blind75

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > 0 {
				end--
			} else if sum < 0 {
				start++
			} else {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start-1] == nums[start] {
					start++
				}
			}
		}
	}

	return res
}
