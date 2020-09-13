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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    root := &TreeNode{preorder[0], nil, nil}
    i, n := 0, len(inorder) - 1
    for ; i <= n; i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1:i+1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root
}
