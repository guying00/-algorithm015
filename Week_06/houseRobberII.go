package Week_06

//时间复杂度O(len(nums))，空间复杂度O(1)
func rob(nums []int) int {
    switch len(nums) {
    case 0:
        return 0
    case 1:
        return nums[0]
    case 2:
        return max(nums[0], nums[1])
    default:
        //第一间和最后一间只能选取一间执行偷窃，所以有两种方案
        //分别计算两种方案的最大值，然后取两者的最大值
        var helper func(nums []int) []int
        helper = func(nums []int) []int {
            dp := make([]int, 2)
            dp[0] = nums[0]
            dp[1] = max(nums[0], nums[1])
            for i := 2; i < len(nums); i++ {
                dp[0], dp[1] = dp[1], max(dp[1], dp[0]+nums[i])
            }
            return dp
        }
        return max(helper(nums[:len(nums)-1])[1], helper(nums[1:])[1])
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
