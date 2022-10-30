package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(values ...int) *TreeNode {
	return insertTreeNodeLevelOrder(values, 0)
}

func insertTreeNodeLevelOrder(values []int, i int) *TreeNode {
	if i >= len(values) || values[i] == -1 {
		return nil
	}
	root := &TreeNode{Val: values[i]}
	root.Left = insertTreeNodeLevelOrder(values, 2*i+1)
	root.Right = insertTreeNodeLevelOrder(values, 2*i+2)
	return root
}
