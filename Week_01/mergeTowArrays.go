package Week_01

func merge(nums1 []int, m int, nums2 []int, n int)  {
    wIdx := m + n - 1
    rIdx1, rIdx2 := m - 1, n - 1
    for rIdx1 >= 0 && rIdx2 >= 0 {
        if nums1[rIdx1] > nums2[rIdx2] {
            nums1[wIdx] = nums1[rIdx1]
            rIdx1--
        } else {
            nums1[wIdx] = nums2[rIdx2]
            rIdx2--
        }
        wIdx--
    }

    for rIdx2 >=0 {
        nums1[wIdx] = nums2[rIdx2]
        rIdx2--
        wIdx--
    }
}
