package blind75

import "github.com/Mirkelor/algo/common"

func cloneGraph(node *common.Node) *common.Node {
	if node == nil {
		return nil
	}
	return dfsCloneGraph(node, map[int]*common.Node{})
}

func dfsCloneGraph(node *common.Node, visited map[int]*common.Node) *common.Node {
	if cloned, ok := visited[node.Val]; ok {
		return cloned
	}
	clone := &common.Node{Val: node.Val, Neighbors: []*common.Node{}}
	visited[node.Val] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfsCloneGraph(neighbor, visited))
	}
	return clone
}
