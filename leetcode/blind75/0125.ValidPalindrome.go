package blind75

import "strings"

func isPalindrome(s string) bool {
	var sb strings.Builder
	s = strings.ToLower(s)
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			sb.WriteRune(c)
		}
	}
	s = sb.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
