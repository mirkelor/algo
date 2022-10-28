package blind75

import "github.com/Mirkelor/algo/common"

func reverseList(head *common.ListNode) *common.ListNode {
	dummy := head
	var prev *common.ListNode = nil

	for dummy != nil {
		next := dummy.Next
		dummy.Next = prev
		prev = dummy
		dummy = next
	}

	return prev
}
