package Week_08

func isPowerOfTwo(n int) bool {
    mid := 1
    for mid <= n {
        if mid == n {
            return true
        }
        mid *=2
    }
    return false
}

func isPowerOfTwo1(n int) bool {
    return n > 0 && n & (n-1) == 0
}
