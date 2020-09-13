package Week_03

//递归解法
//时间复杂度为组合选择的时间复杂度，递归调用层数最大等于节点数，递归栈最大深度k+1，候选解空间k，空间复杂度O(k)
func combine(n int, k int) [][]int {
    result := [][]int{}
    cand := make([]int, k)
    var helper func(curIdx, begin int)
    helper = func(curIdx, begin int) {
        if curIdx == k {
            result = append(result, append([]int(nil), cand...))
            return
        }

        //若剩余数字全选都不能形成 k 个数字的组合，则返回不再处理
        if k-curIdx > n-begin+1 {
            return
        }

        for i := begin; i <= n; i++ {
            cand[curIdx] = i
            helper(curIdx+1, i+1)
        }
    }
    helper(0, 1)
    return result
}
