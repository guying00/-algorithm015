package Week_06

//时间复杂度O(n)，空间复杂度O(n)
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]int, len(nums))
    dp[0] = nums[0]
    maxVal := dp[0]
    for i := 1; i < len(nums); i++ {
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + nums[i]
        } else {
            dp[i] = nums[i]
        }
        if dp[i] > maxVal {
            maxVal = dp[i]
        }
    }
    return maxVal
}

//时间复杂度O(n)，空间复杂度O(1)
func maxSubArray1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    curSum := nums[0]
    maxVal := curSum
    for i := 1; i < len(nums); i++ {
        if curSum > 0 {
            curSum = curSum + nums[i]
        } else {
            curSum = nums[i]
        }
        if curSum > maxVal {
            maxVal = curSum
        }
    }
    return maxVal
}
