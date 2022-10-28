package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(args ...int) *ListNode {
	if len(args) == 0 {
		return new(ListNode)
	}
	root := new(ListNode)
	dummy := root
	for _, arg := range args {
		node := ListNode{Val: arg}
		dummy.Next = &node
		dummy = dummy.Next
	}
	return root.Next
}
