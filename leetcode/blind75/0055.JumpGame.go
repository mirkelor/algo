package blind75

func canJump(nums []int) bool {
	goal := 0
	for i, num := range nums {
		if i > goal {
			return false
		}
		if i+num > goal {
			goal = i + num
		}
	}
	return true
}
