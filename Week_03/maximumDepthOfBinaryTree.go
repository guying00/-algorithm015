package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归解法
//每个节点访问一次，时间复杂度O(n)，递归调用层数最大等于节点数，空间复杂度O(n)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//迭代解法
//每个节点访问一次，时间复杂度O(n)，最多可能需要缓存n/2个节点，空间复杂度O(n)
func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depth := 0
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        depth++
        curLen := len(cache)
        for _, node := range cache {
            if node.Left != nil {
                cache = append(cache, node.Left)
            }
            if node.Right != nil {
                cache = append(cache, node.Right)
            }
        }
        cache = cache[curLen:]
    }
    return depth
}
