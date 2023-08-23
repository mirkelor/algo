package blind75

import "github.com/mirkelor/algo/common"

func reverseList(head *common.ListNode) (prev *common.ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
