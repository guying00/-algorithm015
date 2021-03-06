package Week_09

//时间复杂度O(n)，空间复杂度O(n), n为字符串长度
func longestValidParentheses(s string) int {
    result := 0
    dp := make([]int, len(s))
    for i := 1; i < len(s); i++ {
        if s[i] == ')' {
            if s[i-1] == '(' {
                if i >= 2 {
                    dp[i] = dp[i - 2] + 2
                } else {
                    dp[i] = 2
                }
            } else if i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] == '(' {
                if i - dp[i - 1] >= 2 {
                    dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2
                } else {
                    dp[i] = dp[i - 1] + 2
                }
            }
            result = max(result, dp[i])
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
