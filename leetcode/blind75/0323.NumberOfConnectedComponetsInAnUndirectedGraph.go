package blind75

func countComponents(n int, edges [][]int) int {
	par := make([]int, n)
	for i := range par {
		par[i] = i
	}

	rank := make([]int, n)
	for i := range rank {
		rank[i] = 1
	}

	var find = func(node int) int {
		res := node
		for res != par[res] {
			par[res] = par[par[res]]
			res = par[res]
		}
		return res
	}

	var union = func(node1, node2 int) int {
		p1, p2 := find(node1), find(node2)

		if p1 == p2 {
			return 0
		}

		if rank[p2] > rank[p1] {
			par[p1] = p2
			rank[p2] += rank[p1]
		} else {
			par[p2] = p1
			rank[p1] += rank[p2]
		}

		return 1

	}

	res := n
	for _, edge := range edges {
		res -= union(edge[0], edge[1])
	}
	return res
}
