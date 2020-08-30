package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    rList := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return rList
}

func reverseList1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var newList *ListNode
    ptr1 := head
    ptr2 := ptr1.Next
    for ptr1 != nil {
        ptr1.Next = newList
        newList = ptr1

        ptr1 = ptr2
        if ptr1 != nil {
            ptr2 = ptr1.Next
        }
    }
    return newList
}
