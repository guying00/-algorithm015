package Week_01

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) <= 0 || k <= 0 {
        return nil
    }

    result := []int{}
    n := len(nums)
    max := math.MinInt32
    for i := 0; i < k; i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }
    result = append(result, max)

    for j := k; j < n; j++ {
        if nums[j] > max {
            max = nums[j]
        }
        result = append(result, max)
    }
    return result
}