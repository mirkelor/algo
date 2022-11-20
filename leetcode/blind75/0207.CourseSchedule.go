package blind75

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	m := make(map[int][]int)
	for _, p := range prerequisites {
		m[p[1]] = append(m[p[1]], p[0])
	}

	visit := make([]bool, numCourses)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visit[node] {
			return false
		}
		if len(m[node]) == 0 {
			return true
		}
		visit[node] = true
		for _, neighbor := range m[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		visit[node] = false
		m[node] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
