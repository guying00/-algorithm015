package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归解法。
//每个节点访问一次，时间复杂度O(n)，当树退化为单链表时，栈调用深度为n，空间复杂度O(n)
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    return root
}

//迭代解法。
//每个节点访问一次，时间复杂度O(n)，需要缓存的节点数，等于单层节点最多的一层，最多n/2，依然是n级别，空间复杂度O(n)
func invertTree1(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    cache := []*TreeNode{root}
    for len(cache) > 0 {
        curLen := len(cache)
        for _, curNode := range cache {
            if curNode != nil {
                cache = append(cache, curNode.Left)
                cache = append(cache, curNode.Right)
                curNode.Left, curNode.Right = curNode.Right, curNode.Left
            }
        }
        cache = cache[curLen:]
    }
    return root
}
