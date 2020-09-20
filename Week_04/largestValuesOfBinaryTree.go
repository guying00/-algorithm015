package Week_04

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//dfs解法
//时间复杂度O(n)，空间复杂度O(n)
func largestValues(root *TreeNode) []int {
    result := []int{}
    var helper func(root *TreeNode, level int)
    helper = func(root *TreeNode, level int) {
        if root == nil {
            return
        }

        if len(result) <= level {
            result = append(result, root.Val)
        } else {
            if result[level] < root.Val {
                result[level] = root.Val
            }
        }
        helper(root.Left, level+1)
        helper(root.Right, level+1)
    }
    helper(root, 0)
    return result
}

//bfs解法
//时间复杂度O(n)，空间复杂度O(n)
func largestValues1(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    result := []int{}
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        curLen := len(cache)
        max := math.MinInt64
        for _, node := range cache {
            if node.Val > max {
                max = node.Val
            }
            if node.Left != nil {
                cache = append(cache, node.Left)
            }
            if node.Right != nil {
                cache = append(cache, node.Right)
            }
        }
        result = append(result, max)
        cache = cache[curLen:]
    }
    return result
}
