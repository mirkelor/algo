package blind75

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}

	max := 0
	for num := range set {
		if _, ok := set[num-1]; !ok {
			length := 0
			for _, ok := set[num+length]; ok; {
				length++
				_, ok = set[num+length]
			}
			if length > max {
				max = length
			}
		}
	}

	return max
}
