package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        curNode := cache[len(cache)-1]
        cache = cache[:len(cache)-1]
        if curNode != nil {
            result = append(result, curNode.Val)
            cache = append(cache, curNode.Right)
            cache = append(cache, curNode.Left)
        }
    }
    return result
}
