package Week_03

//时间复杂度O(n*2^n)
//空间复杂度O(n*2^n)
func subsets(nums []int) [][]int {
    return helper(nums, 0)
}

func helper(nums []int, curIdx int) [][]int {
    if len(nums) == curIdx {
        return [][]int{{}}
    }

    result := [][]int{}
    for _, item := range helper(nums, curIdx+1) {
        it := make([]int, len(item))
        copy(it, item)
        result = append(result, it)
        item = append(item, nums[curIdx])
        result = append(result, item)
    }
    return result
}
