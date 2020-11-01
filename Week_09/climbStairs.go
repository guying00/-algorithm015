package Week_09

//时间复杂度O(n)，空间复杂度O(1)
func climbStairs(n int) int {
    if n <= 1 {
        return 1
    }

    dp := make([]int, 2)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        dp[i%2] = dp[0] + dp[1]
    }
    return dp[n%2]
}
