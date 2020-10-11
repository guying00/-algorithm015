package Week_06

//时间复杂度O(n)，空间复杂度O(1)
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    result := nums[0]
    curMaxProd := result
    curMinProd := result
    for i := 1; i < len(nums); i++ {
        maxVal := curMaxProd
        curMaxProd = max(curMaxProd*nums[i], max(nums[i], curMinProd*nums[i]))
        curMinProd = min(curMinProd*nums[i], min(nums[i], maxVal*nums[i]))
        result = max(result, curMaxProd)
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

