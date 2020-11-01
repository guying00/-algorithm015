package Week_09

//时间复杂度O(n)，空间复杂度O(n), n为字符串长度
func reverseStr(s string, k int) string {
    result := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        result[i] = s[i]
    }

    for i := 0; i <= len(s) / k; i++ {
        if i % 2 == 0 {
            m, n := i*k, min(i*k+k-1, len(s)-1)
            for ; m < n; m, n = m+1, n-1 {
                result[m], result[n] = result[n], result[m]
            }
        }
    }
    return string(result)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
