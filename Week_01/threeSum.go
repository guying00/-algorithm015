package Week_01

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := make([][]int, 0, 10)
    lenth := len(nums)
    for i := 0; i < lenth; {
        target := -nums[i]
        maxK := lenth-1
        for j := i+1; j < maxK; {
            for k := maxK; j < k; {
                if nums[j]+nums[k] <= target {
                    if nums[j]+nums[k] == target {
                        result = append(result, []int{nums[i], nums[j], nums[k]})
                    }
                    maxK = k
                    break
                }
                for k--; k > j && (nums[k] == nums[k+1] || nums[j]+nums[k] > target); k-- {
                    maxK = k
                }
            }
            for j++; j < maxK && nums[j] == nums[j-1]; j++ {}
        }
        for i++; i < len(nums) && nums[i] == nums[i-1]; i++ {}
    }
    return result
}
