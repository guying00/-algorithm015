package Week_06

//时间复杂度O(len(text1)*len(text1))，空间复杂度O(min(len(text1), len(text1)))
func longestCommonSubsequence(text1 string, text2 string) int {
    if text1 == "" || text2 == "" {
        return 0
    }

    if len(text1) > len(text2) {
        text1, text2 = text2, text1
    }

    n := len(text1)+1
    m := len(text2)+1
    dp := make([][]int, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([]int, m)
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i%2][j] = dp[(i-1)%2][j-1] + 1
            } else {
                dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
            }
        }
    }
    return dp[(n-1)%2][len(dp[0])-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
