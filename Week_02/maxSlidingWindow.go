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

//滑动窗口处理，要做的其实就是3件事情：
//1、每次加入元素到窗口时，将前面比他小的元素都提出队列，保持队列中元素的递减序
//2、每次加入元素后，检查对头元素是否已滑出窗口，已滑出则删除队头元素
//3、从窗口中有k个元素开始，不停输出队头元素即可
func maxSlidingWindow1(nums []int, k int) []int {
    result := []int{}
    queue := []int{}
    for i, n := range nums {
        j := len(queue)-1
        for ; j >= 0 && nums[queue[j]] <= n; j-- {}
        queue = queue[:j+1]

        queue = append(queue, i)
        if i - queue[0] >= k {
            queue = queue[1:]
        }

        if i >= k-1 {
            result = append(result, nums[queue[0]])
        }
    }
    return result
}
