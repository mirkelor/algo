package common

type Node struct {
	Val       int
	Neighbors []*Node
}

func NewNode(adjList [][]int) *Node {
	m := make(map[int]*Node)

	for i := range adjList {
		m[i+1] = &Node{Val: i + 1, Neighbors: []*Node{}}
	}

	for i, adj := range adjList {
		for _, neighbor := range adj {
			m[i+1].Neighbors = append(m[i+1].Neighbors, m[neighbor])
		}
	}

	return m[1]
}
