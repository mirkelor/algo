package blind75

import "github.com/mirkelor/algo/common"

func hasCycle(head *common.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if fast != nil && slow == fast {
			return true
		}
	}
	return false
}
