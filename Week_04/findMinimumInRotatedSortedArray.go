package Week_04

func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    mid := (left+right) / 2
    for left <= right {
        if mid > 0 && nums[mid-1] > nums[mid] {
            return nums[mid]
        }
        if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
            return nums[mid+1]
        }
        //数组左半部份有序
        if nums[0] <= nums[mid]  {
            //目标值在左半部份
            left = mid+1
        } else {  //数组右半部份有序
            //目标值在左半部份
            right = mid-1
        }
        mid = (left+right) / 2
    }
    return nums[0]
}
