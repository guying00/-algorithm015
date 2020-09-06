package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {
        return nil
    }

    result := [][]int{}
    cache := []*Node{root}
    for len(cache) > 0 {
        curLen := len(cache)
        levelVal := []int{}
        for _, curNode := range cache {
            levelVal = append(levelVal, curNode.Val)
            cache = append(cache, curNode.Children...)
        }
        result = append(result, levelVal)
        cache = cache[curLen:]
    }
    return result
}
