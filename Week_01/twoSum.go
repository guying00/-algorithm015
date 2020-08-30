package Week_01

func twoSum(nums []int, target int) []int {
    cache := map[int]int{}
    for i, elem := range nums {
        idx, ok := cache[elem]
        if ok {
            return []int{idx, i}
        }
        cache[target-elem] = i
    }
    return nil
}
