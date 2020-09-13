package Week_03

//时间复杂度O(n), 空间复杂度O(1)
func climbStairs(n int) int {
    cache0, cache1 := 1, 1
    for i := 2; i <= n; i++ {
        cache0, cache1 = cache1, cache0+cache1
    }
    return cache1
}

//时间复杂度O(n), 空间复杂度O(1)
func climbStairs1(n int) int {
    cache := []int{1, 1}
    i := 2
    for ; i <= n; i++ {
        cache[i%2] = cache[0] + cache[1]
    }
    return cache[(i-1)%2]
}

//时间复杂度O(n), 空间复杂度O(n)
var cache = map[int]int{0: 1, 1: 1}
func climbStairs2(n int) int {
    if v, ok := cache[n]; ok {
        return v
    }
    cache[n] = climbStairs(n-1) + climbStairs(n-2)
    return cache[n]
}
