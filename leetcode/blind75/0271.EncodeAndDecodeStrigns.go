package blind75

import (
	"fmt"
	"strconv"
)

func encode(strs []string) string {
	res := ""
	for _, s := range strs {
		res += fmt.Sprintf("%v#%s", len(s), s)
	}
	return res
}

func decode(str string) []string {
	res, i := make([]string, 0), 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		l, _ := strconv.Atoi(str[i:j])
		res = append(res, str[j+1:j+1+l])
		i = j + 1 + l
	}
	return res
}
