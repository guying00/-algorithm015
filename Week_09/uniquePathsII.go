package Week_09

//时间复杂度O(len(obstacleGrid)*len(obstacleGrid[0]))，空间复杂度O(len(obstacleGrid[0]))
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
        return 0
    }

    dp := make([]int, len(obstacleGrid[0]))
    for j := 0; j < len(obstacleGrid[0]); j++ {
        if obstacleGrid[0][j] == 1 {
            break
        }
        dp[j] = 1
    }

    for i := 1; i < len(obstacleGrid); i++ {
        if obstacleGrid[i][0] == 1 {
            dp[0] = 0
        }

        for j := 1; j < len(obstacleGrid[0]); j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
            } else {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[len(dp)-1]
}
