package Week_08

func reverseBits(num uint32) uint32 {
    res := uint32(0)
    max := uint32(1 << 31)
    for i := 0; i < 32; i++ {
        if num & (1 << i) > 0 {
            res |= (max >> i)
        }
    }
    return res
}

