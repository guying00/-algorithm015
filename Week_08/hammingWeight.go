package Week_08

func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        if num & uint32(1) > 0 {
            count++
        }
        num = num >> 1
    }
    return count
}
