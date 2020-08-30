package Week_01

func removeDuplicates(nums []int) int {
    widx := 0
    for ridx := 0; ridx < len(nums); ridx++ {
        if ridx < len(nums) - 1 && nums[ridx] == nums[ridx + 1] {
            continue
        }

        nums[widx] = nums[ridx]
        widx++
    }
    return widx
}