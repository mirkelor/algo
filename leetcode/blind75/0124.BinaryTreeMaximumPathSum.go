package blind75

import "github.com/mirkelor/algo/common"

func maxPathSum(root *common.TreeNode) int {
	res := -1 << 31
	maxPath(root, &res)
	return res
}

func maxPath(root *common.TreeNode, res *int) int {
	if root == nil {
		return -1 << 31
	}
	l := maxPath(root.Left, res)
	r := maxPath(root.Right, res)
	m := root.Val + max(r, l, 0)
	*res = max(*res, m, root.Val+r+l)
	return m
}

func max(a ...int) int {
	m := a[0]
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}
