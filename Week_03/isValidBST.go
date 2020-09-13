package Week_03

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//迭代解法。
//每个节点访问一次，时间复杂度O(n)，需要缓存的节点数，等于单层节点最多的一层，最多n/2，依然是n级别，空间复杂度O(n)
func isValidBST(root *TreeNode) bool {
    var preNode *TreeNode
    cache := []*TreeNode{}
    for len(cache) > 0 || root != nil {
        for root != nil {
            cache = append(cache, root)
            root = root.Left
        }

        root = cache[len(cache)-1]
        cache = cache[:len(cache)-1]
        if preNode != nil && preNode.Val >= root.Val {
            return false
        }
        preNode = root
        root = root.Right
    }
    return true
}

//递归解法，根据访问路径的左右子树，确定各子树节点值的上下界
//每个节点访问一次，时间复杂度O(n)，空间复杂度O(n)
//需要注意的一点是，当节点值为int可表示的最小值或最大值时，若树的左或右子树为空，搜索树是有效的。
func isValidBST1(root *TreeNode) bool {
    var helper func(root *TreeNode, down, up int) bool
    helper = func(root *TreeNode, down, up int) bool {
        if root == nil {
            return true
        }
        return (down < root.Val || math.MinInt64 == root.Val && root.Left == nil) &&
            (root.Val < up || math.MaxInt64 == root.Val && root.Right == nil) &&
            helper(root.Left, down, min(root.Val, up)) &&
            helper(root.Right, max(root.Val, down), up)
    }
    return helper(root, math.MinInt64, math.MaxInt64)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
