package blind75

import "github.com/mirkelor/algo/common"

func kthSmallest(root *common.TreeNode, k int) int {
	res := -1
	count := 0
	var dfs func(*common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if count++; count == k {
			res = node.Val
			return
		} else if count > k {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
