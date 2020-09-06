package Week_02

import "sort"

func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    result := []int{}
    for i := 0; i < k; i++ {
        result = append(result, arr[i])
    }
    return result
}
