package Week_09

//时间复杂度O(len(nums))，空间复杂度O(1)
func rob(nums []int) int {
    switch len(nums) {
    case 0:
        return 0
    case 1:
        return nums[0]
    default:
        dp := make([]int, 2)
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])
        for i := 2; i < len(nums); i++ {
            dp[0], dp[1] = dp[1], max(dp[1], dp[0]+nums[i])
        }
        return dp[1]
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
