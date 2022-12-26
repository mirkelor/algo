package blind75

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var backtrack func(idx, sum int, comb []int)
	backtrack = func(idx, sum int, comb []int) {
		if sum > target {
			return
		} else if sum == target {
			res = append(res, append([]int{}, comb...))
		} else {
			for i := idx; i < len(candidates); i++ {
				comb = append(comb, candidates[i])
				backtrack(i, sum+candidates[i], comb)
				comb = comb[:len(comb)-1]
			}
		}
	}
	backtrack(0, 0, []int{})
	return res
}
