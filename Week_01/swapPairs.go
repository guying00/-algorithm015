package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := head.Next
    for head != nil && head.Next != nil {
        remain := head.Next.Next
        head.Next.Next = head
        if remain != nil && remain.Next != nil {
            head.Next = remain.Next
        } else {
            head.Next = remain
        }
        head = remain
    }
    return newHead
}

func swapPairs1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head
    return newHead
}
