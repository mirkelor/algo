package blind75

import "github.com/Mirkelor/algo/common"

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
    root := new(common.ListNode)
    tail := root
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }
    if list1 != nil {
        tail.Next = list1
    }
    if list2 != nil {
        tail.Next = list2
    }
    return root.Next
}