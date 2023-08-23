package blind75

import "github.com/mirkelor/algo/common"

func reorderList(head *common.ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *common.ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	tail := head
	for prev != nil && prev.Next != nil {
		tail.Next, tail = prev, tail.Next
		prev.Next, prev = tail, prev.Next
	}

}
