package Week_03

import "sort"

//同层数据回溯迭代，不同层的数据递归处理，对同一层的元素，先后将其放入结果集中，然后将剩余元素组成新的列表递归到下一层处理
//直到缓冲区填满时，所有元素即处理完毕，输出本次结果，返回上层回溯状态继续处理
//时间复杂度O(n!×n), 空间复杂度O(n)
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    cand := make([]int, len(nums))
    var helper func(curIdx int, remain []int)
    helper = func(curIdx int, remain []int) {
        if curIdx == len(nums) {
            result = append(result, append([]int(nil), cand...))
            return
        }

        for i, n := range remain {
            if i > 0 && remain[i-1] == n {
                continue
            }
            cand[curIdx] = n
            helper(curIdx+1, append(append([]int(nil), remain[:i]...), remain[i+1:]...))
        }
    }
    helper(0, nums)
    return result
}
