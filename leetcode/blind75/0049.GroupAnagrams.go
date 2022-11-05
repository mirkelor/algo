package blind75

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sorted := sortAnagramString(str)
		m[sorted] = append(m[sorted], str)
	}

	res := make([][]string, 0, len(m))
	for _, strs := range m {
		res = append(res, strs)
	}
	return res
}

func sortAnagramString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool { return runes[i] > runes[j] })
	return string(runes)
}
