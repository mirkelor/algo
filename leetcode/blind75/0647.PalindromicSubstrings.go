package blind75

func countSubstrings(s string) int {
	sum := 0
	for i := range s {
		sum += countSubstringsHelper(s, i, i)
		sum += countSubstringsHelper(s, i, i+1)
	}
	return sum
}

func countSubstringsHelper(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}
