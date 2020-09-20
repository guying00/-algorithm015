package Week_04

//时间复杂度O(n)，空间复杂度O(1)
func maxProfit(prices []int) int {
    profit := 0
    for i := 0; i < len(prices); i++ {
        if i > 0 && prices[i] > prices[i-1] {
            profit += (prices[i] - prices[i-1])
        }
    }
    return profit
}
