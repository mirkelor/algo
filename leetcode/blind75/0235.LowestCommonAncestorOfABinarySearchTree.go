package blind75

import "github.com/Mirkelor/algo/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
