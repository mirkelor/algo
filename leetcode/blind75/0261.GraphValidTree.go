package blind75

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visit := make(map[int]struct{})

	var dfs func(i, prev int) bool
	dfs = func(i, prev int) bool {
		if _, ok := visit[i]; ok {
			return false
		}
		visit[i] = struct{}{}
		for _, j := range adj[i] {
			if j == prev {
				continue
			}
			if !dfs(j, i) {
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && n == len(visit)
}
