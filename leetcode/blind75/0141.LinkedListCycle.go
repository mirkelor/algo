package blind75

import "github.com/Mirkelor/algo/common"

func hasCycle(head *common.ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast != nil && slow == fast {
            return true
        }
    }
    return false
}
