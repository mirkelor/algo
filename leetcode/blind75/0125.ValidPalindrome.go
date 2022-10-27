package blind75

import "strings"

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if !isValidPalindrome(s[start]) {
			start++
			continue
		}
		if !isValidPalindrome(s[end]) {
			end--
			continue
		}
		if !strings.EqualFold(string(s[start]), string(s[end])) {
			return false
		}
		start++
		end--
	}
	return true
}

func isValidPalindrome(c byte) bool {
	return ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9')
}