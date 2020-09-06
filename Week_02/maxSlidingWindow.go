package Week_02

func maxSlidingWindow(nums []int, k int) []int {
    result := []int{}
    queue := make([]int, 0, k)
    for i, n := range nums {
        if len(queue) > 0 && i - queue[0] >= k {
            queue = queue[1:]
        }

        j := len(queue) - 1
        for j >= 0 && nums[queue[j]] <= n {
            j--
        }
        queue = queue[:j+1]
        queue = append(queue, i)

        if i >= k-1 {
            result = append(result, nums[queue[0]])
        }
    }
    return result
}
