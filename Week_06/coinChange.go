package Week_06

//时间复杂度O(len(coins)*amount)，空间复杂度O(amount)
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount+1
    }

    for _, coin := range coins {
        for am := 0; am <= amount; am++ {
            if am - coin >= 0 {
                dp[am] = min(dp[am-coin]+1, dp[am])
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
