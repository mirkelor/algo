package blind75

import (
	"sort"

	"github.com/Mirkelor/algo/common"
)

func alienOrder(words []string) string {
	adj := make(map[byte][]byte)

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := common.Min(len(w1), len(w2))
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				adj[w1[j]] = append(adj[w1[j]], w2[j])
				break
			}
		}
	}

	// false=visited, true=current path
	visit := make(map[byte]bool)
	res := make([]byte, 0)

	var dfs func(c byte) bool
	dfs = func(c byte) bool {
		if _, ok := visit[c]; ok {
			return visit[c]
		}

		visit[c] = true

		for _, nei := range adj[c] {
			if dfs(nei) {
				return true
			}
		}

		visit[c] = false
		res = append(res, c)
		return false
	}

	for c := range adj {
		if dfs(c) {
			return ""
		}
	}

	sort.Slice(res, func(i, j int) bool { return i > j })

	return string(res)
}
