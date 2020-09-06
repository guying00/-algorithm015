package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }

    result := []int{}
    cache := []*Node{root}
    for len(cache) > 0 {
        curNode := cache[len(cache)-1]
        cache = cache[:len(cache)-1]
        result = append(result, curNode.Val)
        for i := len(curNode.Children) - 1; i >= 0; i-- {
            cache = append(cache, curNode.Children[i])
        }
    }
    return result
}
