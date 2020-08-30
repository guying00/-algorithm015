package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for {
        if fast == nil || fast.Next == nil {
            return nil
        }
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }

    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}

func detectCycle1(head *ListNode) *ListNode {
    cache := map[*ListNode]bool{}
    for ptr := head; ptr != nil; ptr = ptr.Next {
        if cache[ptr] {
            return ptr
        }
        cache[ptr] = true
    }
    return nil
}
