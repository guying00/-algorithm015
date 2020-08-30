package Week_01

func moveZeroes(nums []int)  {
    wptr := 0
    for rptr := 0; rptr < len(nums); rptr++ {
        if nums[rptr] != 0 {
            if wptr != rptr {
                nums[wptr] = nums[rptr]
                nums[rptr] = 0
            }
            wptr++
        }
    }
}