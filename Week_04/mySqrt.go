package Week_04

import "math"

//采用牛顿迭代法
//时间复杂度O(logn)，空间复杂度O(1)
func mySqrt(x int) int {
    if x <= 0 {
        return 0
    }
    c, x0 := float64(x), float64(x)
    for {
        xi := 0.5 * (x0+c/x0)
        if math.Abs(xi - x0) < 1e-7 {
            break
        }
        x0 = xi
    }
    return int(x0)
}

//采用二分查找
//时间复杂度O(logn)，空间复杂度O(1)
func mySqrt1(x int) int {
    ans := 0
    l, r := 0, x
    for l <= r {
        mid := (l+r) / 2
        if mid*mid <= x {
            ans = mid
            l = mid+1
        } else {
            r = mid-1
        }
    }
    return ans
}