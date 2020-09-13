package Week_03

import "math/rand"

//投票法，时间复杂度O(n) 空间复杂度O(1)
func majorityElement(nums []int) int {
    cand, count := 0, 0
    for _, n := range nums {
        if count == 0 {
            cand = n
        }
        if cand == n {
            count++
        } else {
            count--
        }
    }
    return cand
}

//排序法，时空复杂度同使用的排序算法
func majorityElement3(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

//哈希表统计，时间复杂度O(n)，空间复杂度O(n)
func majorityElement4(nums []int) int {
    statis := map[int]int{}
    cand, max := 0, 0
    for _, n := range nums {
        statis[n]++
        if statis[n] > max {
            cand = n
            max = statis[n]
        }
    }
    return cand
}

//随机法，最坏时间复杂度O(+Infinity), 最好时间复杂度O(n) 空间复杂度O(1)
func majorityElement1(nums []int) int {
    for {
        cand := nums[rand.Intn(len(nums))]
        count := 0
        for _, n := range nums {
            if n == cand {
                count++
            }
            if count > len(nums) / 2 {
                return cand
            }
        }
    }
    return 0
}

//分治法，时间复杂度O(nlogn)，空间复杂度O(logn)
func majorityElement2(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    cand1, cand2 := majorityElement(nums[:len(nums)/2]), majorityElement(nums[len(nums)/2:])
    if cand1 != cand2 {
        count := 0
        for _, n := range nums {
            if cand1 == n {
                count++
            }
            if count > len(nums) / 2 {
                return cand1
            }
        }
        return cand2
    }
    return cand1
}
