package Week_09

//时间复杂度O(n^2)，空间复杂度O(1), n为字符串长度
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    maxLen := 1
    dp := make([]int, len(nums))
    dp[0] = 1
    for i := 1; i < len(nums); i++ {
        lenth := 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] && dp[j] > lenth {
                lenth = dp[j]
            }
        }
        dp[i] = lenth+1
        if maxLen < dp[i] {
            maxLen = dp[i]
        }
    }
    return maxLen
}
