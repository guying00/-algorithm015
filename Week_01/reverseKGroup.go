package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    last := head
    count := 1
    for last != nil && count < k {
        last = last.Next
        count++
    }

    if last == nil || last == head {
        return head
    }

    //翻转前 k - 1 个节点，preHead中保存翻转后的局部链表头
    preHead, curNode := head, head.Next
    for ; curNode != last; {
        next := curNode.Next
        curNode.Next = preHead
        preHead = curNode
        curNode = next
    }
    //将原链表的第一个节点链接到剩余列表翻转后的链表前面
    head.Next = reverseKGroup(last.Next, k)
    //将当前层的第 k 个节点链接到翻转后的链表头部，作为新的链表头
    last.Next = preHead
    return last
}

func reverseLinkedList(head, last *ListNode) (*ListNode, *ListNode) {
    preHead, curNode := head, head.Next
    for ; curNode != last; {
        next := curNode.Next
        curNode.Next = preHead
        preHead = curNode
        curNode = next
    }
    last.Next = preHead
    return last, head
}