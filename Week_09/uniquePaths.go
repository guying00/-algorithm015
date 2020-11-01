package Week_09

//时间复杂度O(m*n)，空间复杂度O(min(m, n))
func uniquePaths(m int, n int) int {
    if m < 1 || n < 1 {
        return 0
    }

    if m > n {
        m, n = n, m
    }

    dp := make([]int, m)
    for i := 0; i < m; i++ {
        dp[i] = 1
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[j] += dp[j-1]
        }
    }
    return dp[m-1]
}
