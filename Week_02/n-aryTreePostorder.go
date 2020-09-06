package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return nil
    }

    result := []int{}
    cache := []*Node{root}
    for len(cache) > 0 {
        curNode := cache[len(cache)-1]
        if curNode == nil {
            curNode = cache[len(cache)-2]
            cache = cache[:len(cache)-2]
            result = append(result, curNode.Val)
        } else {
            cache = append(cache, nil)  //插入空指针，用于标记当前处理的节点已经扩展访问，当弹出空指针时，位于其前一位置的节点可以访问并出栈
            for i := len(curNode.Children) - 1; i >= 0; i-- {
                cache = append(cache, curNode.Children[i])
            }
        }
    }
    return result
}
