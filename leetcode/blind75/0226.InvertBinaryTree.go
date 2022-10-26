package blind75

import "github.com/Mirkelor/algo/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	temp := invertTree(root.Right)
	root.Right = invertTree(root.Left)
	root.Left = temp
	return root
}
