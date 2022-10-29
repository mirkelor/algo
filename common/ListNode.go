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

func NewListNodeWithCycle(pos int, args ...int) *ListNode {
	if len(args) == 0 {
		return new(ListNode)
	}
	if pos < 0 || pos >= len(args) {
		return NewListNode(args...)
	}

	root := new(ListNode)
	dummy := root
	for i := 0; i < pos; i++ {
		node := ListNode{Val: args[i]}
		dummy.Next = &node
		dummy = dummy.Next
	}

	cycle := dummy
	for i := pos; i < len(args); i++ {
		node := ListNode{Val: args[i]}
		dummy.Next = &node
		dummy = dummy.Next
	}
	dummy.Next = cycle
	return root.Next
}
