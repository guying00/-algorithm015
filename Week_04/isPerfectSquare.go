package Week_04

//采用二分查找
//时间复杂度O(logn)，空间复杂度O(1)
func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }

    left, right := 0, num/2
    cand := (left+right)/2
    for left <= right {
        if cand*cand > num {
            right = cand-1
        } else {
            left = cand+1
        }
        cand = (left+right)/2
    }
    return cand*cand == num
}
