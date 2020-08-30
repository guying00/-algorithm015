package Week_01

func climbStairs(n int) int {
    var cache = make([]int, 2)
    cache[0] = 2
    cache[1] = 1
    for i := 3; i <= n; i++ {
        cache[i % 2] = cache[0] + cache[1]
    }
    return cache[n % 2]
}
