package blind75

func longestPalindrome(s string) string {
	l := 0
	r := 0
	slen := len(s)

	for i := 0; i < slen; i++ {
		l1, r1 := longestPalindromeHelper(s, i, i)
		if r1-l1-1 > r-l {
			l = l1 + 1
			r = r1
		}

		l2, r2 := longestPalindromeHelper(s, i, i+1)
		if r2-l2-1 > r-l {
			l = l2 + 1
			r = r2
		}
	}

	return s[l:r]
}

func longestPalindromeHelper(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l, r
}
