package Week_03

import (
    "strconv"
    "strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//二叉树的序列化与反序列化，在技术上就是将在二叉树遍历（不管前中后序）的过程中，将节点的值进行序列化保存，并使用空来标记二叉树的逻辑结构（为空的子树）；在反序列化的时候，
//按相同的逻辑去理解字符串，就可以还原原先的数据和树结构。
//以下程序按前序遍历的方式保存和反序列化二叉树，序列化时，每个节点访问1次，时间复杂度 O(n)，空间复杂度 O(n)；反序列化时，生成节点结构一次，重建逻辑结构1次，每个节点访问2次，
//时间复杂度依然是 O(n)，使用的空间也是 O(n)
type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    result := []byte{'['}
    lastNotNull := 5
    cache := []*TreeNode{root}
    for len(cache) > 0 {
        curLen := len(cache)
        for _, curNode := range cache {
            if curNode == nil {
                result = append(result, []byte("null")...)
            } else {
                result = append(result, []byte(strconv.Itoa(curNode.Val))...)
                lastNotNull = len(result)
                cache = append(cache, curNode.Left)
                cache = append(cache, curNode.Right)
            }
            result = append(result, ',')
        }
        cache = cache[curLen:]
    }
    result = result[:lastNotNull+1]
    result[len(result)-1] = ']'
    return string(result)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    d := data[1:len(data)-1]
    nodeVals := strings.Split(d, ",")
    cache := make([]*TreeNode, len(nodeVals))
    for i, val := range nodeVals {
        if val != "null" {
            if v, err := strconv.Atoi(val); err == nil {
                cache[i] = &TreeNode{v, nil, nil}
            }
        }
    }

    getChild := func(idx *int) *TreeNode {
        if *idx < len(cache) {
            *idx++
            return cache[*idx-1]
        }
        return nil
    }
    for i, j := 0, 1; i < len(cache); i++ {
        if cache[i] != nil {
            cache[i].Left = getChild(&j)
            cache[i].Right = getChild(&j)
        }
    }
    return cache[0]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
