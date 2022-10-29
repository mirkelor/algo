package blind75

import "github.com/Mirkelor/algo/common"

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	dummy := &common.ListNode{Next: head}

	slow, fast := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
