package blind75

import "github.com/Mirkelor/algo/common"

func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := make([]*common.TreeNode, 0)
	q = append(q, root)

	res := make([][]int, 0)
	for len(q) > 0 {
		curLen := len(q)
		nodes := make([]int, 0)

		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			nodes = append(nodes, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, nodes)
	}
	return res
}
