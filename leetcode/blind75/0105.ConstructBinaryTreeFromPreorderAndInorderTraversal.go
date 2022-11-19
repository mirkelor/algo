package blind75

import "github.com/Mirkelor/algo/common"

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := new(common.TreeNode)
	root.Val = preorder[0]
	mid := common.FindIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
