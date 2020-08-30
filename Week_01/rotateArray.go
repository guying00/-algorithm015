package Week_01

func rotate(nums []int, k int) {
    if len(nums) <= 1 || k <= 0{
        return
    }

    k = k % len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])

}

func reverse(nums []int) {
    for i, j := 0, len(nums)-1; i < j; {
        nums[i] ^= nums[j]
        nums[j] ^= nums[i]
        nums[i] ^= nums[j]
        i++;j--
    }
}

func rotate2(nums []int, k int) {
    lenth := len(nums)
    k = k % lenth
    if lenth <= 1 || k <= 0 || lenth == k {
        return
    }

    for i := 0; i < gcd(lenth, k); i++ {
        dealOneCycle(nums, i, k, lenth)
    }
}

func dealOneCycle(nums []int, begin int, k int, n int) {
    empty := begin
    tmp := nums[begin]
    for empty != k+begin {
        curIdx := (empty - k + n) % n
        nums[empty] = nums[curIdx]
        empty = curIdx
    }
    nums[empty] = tmp
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func rotate3(nums []int, k int)  {
    lenth := len(nums)
    for i := 1; i <= k; i++ {
        tmp := nums[lenth-1]
        for j := lenth - 2; j >= 0; j-- {
            nums[j+1] = nums[j]
        }
        nums[0] = tmp
    }
}
