package Week_09

//时间复杂度O(n+m)，空间复杂度O(1), n, m为字符串长度
func numJewelsInStones(J string, S string) int {
    if len(J) == 0 || len(S) == 0 {
        return 0
    }

    cache := make([]bool, 128)
    for i := 0; i < len(J); i++ {
        cache[J[i]] = true
    }

    count := 0
    for i := 0; i < len(S); i++ {
        if cache[S[i]] {
            count++
        }
    }
    return count
}
