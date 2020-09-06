package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    cache := []*TreeNode{}
    for len(cache) > 0 || root != nil {
        for root != nil {
            cache = append(cache, root)
            root = root.Left
        }

        root = cache[len(cache)-1]
        cache = cache[:len(cache)-1]
        result = append(result, root.Val)
        root = root.Right
    }
    return result
}
