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
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    switch {
    case root.Left == nil && root.Right == nil:
        return 1
    case root.Left != nil && root.Right != nil:
        return min(minDepth(root.Left), minDepth(root.Right)) + 1
    case root.Left != nil:
        return minDepth(root.Left) + 1
    case root.Right != nil:
        return minDepth(root.Right) + 1
    }
    return 0
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

//迭代解法
//每个节点访问一次，时间复杂度O(n)，最多可能需要缓存n/2个节点，空间复杂度O(n)
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depth := 0
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        depth++
        curLen := len(cache)
        for _, node := range cache {
            if node.Left == nil && node.Right == nil {
                return depth
            }
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
