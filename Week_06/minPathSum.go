package Week_06

//时间复杂度O(len(grid)*len(grid[0]))，空间复杂度O(len(grid[0]))
func minPathSum(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    dp := make([]int, len(grid[0]))
    dp[0] = grid[0][0]
    for j := 1; j < len(grid[0]); j++ {
        dp[j] += dp[j-1] + grid[0][j]
    }

    for i := 1; i < len(grid); i++ {
        dp[0] += grid[i][0]
        for j := 1; j < len(grid[0]); j++ {
            dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
        }
    }
    return dp[len(dp)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
