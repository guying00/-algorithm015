package Week_03

//时间复杂度O(logn)
//空间复杂度O(logn)
func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1 / helper(x, -n)
    }
    return helper(x, n)
}

func helper(x float64, n int) float64 {
    result := 1.0
    for mid := x; n > 0; n = (n >> 1) {
        if n & 1 == 1 {
            result *= mid
        }
        mid *= mid
    }
    return result
}
