package blind75

import "github.com/Mirkelor/algo/common"

func reverseList(head *common.ListNode) (prev *common.ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
