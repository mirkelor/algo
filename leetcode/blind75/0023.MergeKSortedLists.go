package blind75

import "github.com/mirkelor/algo/common"

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := make([]*common.ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeList(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, mergeList(lists[i], nil))
			}
		}
		lists = mergedLists
	}

	return lists[0]
}

func mergeList(l1, l2 *common.ListNode) *common.ListNode {
	merged := new(common.ListNode)
	tail := merged
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	}
	if l2 == nil {
		tail.Next = l1
	}

	return merged.Next
}
