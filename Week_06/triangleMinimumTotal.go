package Week_06

//时间复杂度O(n*n)，空间复杂度O(n)
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
        return 0
    }

    dp := make([]int, len(triangle))
    dp[0] = triangle[0][0]
    result := dp[0]
    for i := 1; i < len(triangle); i++ {
        dp[i] = dp[i-1] + triangle[i][i]
        result = dp[i]
        for j := i-1; j >= 1; j-- {
            dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
            result = min(result, dp[j])
        }
        dp[0] += triangle[i][0]
        result = min(result, dp[0])
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
