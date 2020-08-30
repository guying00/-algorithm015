package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle1(head *ListNode) bool {
    ptr1, ptr2 := head, head
    for ptr1 != nil && ptr2 != nil {
        ptr1 = ptr1.Next
        ptr2 = ptr2.Next
        if ptr2 != nil {
            ptr2 = ptr2.Next
        }
        if ptr1 != nil && ptr1 == ptr2 {
            return true
        }
    }
    return false
}

func hasCycle(head *ListNode) bool {
    for cache := map[*ListNode]bool{}; head != nil; head = head.Next {
        if cache[head] {
            return true
        }
        cache[head] = true
    }
    return false
}