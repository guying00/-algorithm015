package Week_04

//采用二分查找
//时间复杂度O(logn)，空间复杂度O(1)
func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := (left+right) / 2
        if nums[mid] == target {
            return mid
        }

        //数组左半部份有序
        if nums[0] <= nums[mid]  {
            //目标值在左半部份
            if nums[0] <= target && target < nums[mid] {
                right = mid-1
            } else { //目标值在右半部份
                left = mid+1
            }
        } else {  //数组右半部份有序
            //目标值在左半部份
            if nums[mid] < target && target <= nums[len(nums)-1] {
                left = mid+1
            } else { //目标值在右半部份
                right = mid-1
            }
        }
    }
    return -1
}
