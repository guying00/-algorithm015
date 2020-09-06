package Week_02

func twoSum(nums []int, target int) []int {
    cache := map[int]int{}
    for i, n := range nums {
        if idx, ok := cache[n]; ok {
            return []int{idx, i}
        }
        cache[target-n] = i
    }
    return nil
}
