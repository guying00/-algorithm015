package Week_09

//时间复杂度O(n)，空间复杂度O(1), n为楼梯长度
func minCostClimbingStairs(cost []int) int {
    minCost := 0
    cache0, cache1 := min(0, cost[0]), min(cost[0], cost[1])
    for i := 2; i < len(cost); i++ {
        //cache0始终存储(i - 2)的最小cost，cache1始终存储(i - 1)的最小cost
        minCost = min(cache0 + cost[i - 1], cache1 + cost[i])
        cache0, cache1 = cache1, minCost
    }
    return minCost
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
