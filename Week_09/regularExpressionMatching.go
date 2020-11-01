package Week_09

//时间复杂度O(m*n)，空间复杂度O(m*n), m, n为字符串长度
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    matches := func(i, j int) bool {
        if i == 0 {
            return false
        }
        if p[j-1] == '.' {
            return true
        }
        return s[i-1] == p[j-1]
    }

    dp := make([][]bool, m + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, n + 1)
    }
    dp[0][0] = true
    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j] || dp[i][j-2]
                if matches(i, j - 1) {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else if matches(i, j) {
                dp[i][j] = dp[i][j] || dp[i-1][j-1]
            }
        }
    }
    return dp[m][n]
}
