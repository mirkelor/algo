package blind75

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[target-nums[i]]; ok {
			return []int{i, numsMap[target-nums[i]]}
		}
		numsMap[nums[i]] = i
	}
	return nil
}
