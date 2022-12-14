package blind75

func minWindow(s string, t string) string {
	if t == "" || len(s) < len(t) {
		return ""
	}

	window, countT := make(map[byte]int), make(map[byte]int)
	for i := range t {
		countT[t[i]]++
	}

	have, need, res, l := 0, len(countT), []int{-1, 1<<31 - 1}, 0
	for r := range s {
		c := s[r]
		window[c]++

		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		for have == need {
			if r-l+1 < res[1]-res[0]+1 {
				res = []int{l, r}
			}
			window[s[l]]--
			if countT[s[l]] > 0 && window[s[l]] < countT[s[l]] {
				have--
			}
			l++
		}
	}
	if res[1] == 1<<31-1 {
		return s
	}
	return s[res[0] : res[1]+1]
}
