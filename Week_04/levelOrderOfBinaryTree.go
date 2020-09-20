package Week_04

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//广度优先遍历
//时间复杂度：O(N)
//空间复杂度：O(N)
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    result := [][]int{}
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        curLen := len(cache)
        level := []int{}
        for _, node := range cache {
            level = append(level, node.Val)
            if node.Left != nil {
                cache = append(cache, node.Left)
            }
            if node.Right != nil {
                cache = append(cache, node.Right)
            }
        }
        result = append(result, level)
        cache = cache[curLen:]
    }
    return result
}

//深度优先遍历
//时间复杂度：O(N)
//空间复杂度：O(N)
func levelOrder1(root *TreeNode) [][]int {
    result := [][]int{}
    var helper func(root *TreeNode, level int)
    helper = func(root *TreeNode, level int) {
        if root == nil {
            return
        }

        if len(result) <= level {
            result = append(result, []int(nil))
        }
        result[level] = append(result[level], root.Val)
        helper(root.Left, level+1)
        helper(root.Right, level+1)
    }
    helper(root, 0)
    return result
}
