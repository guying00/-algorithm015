package Week_04

//采用贪心法，每次找零时，尽量保存好手里的5美元钞，以达到最大的兑换灵活性
//时间复杂度O(n)，空间复杂度O(1)
func lemonadeChange(bills []int) bool {
    count5, count10 := 0, 0
    for _, bill := range bills {
        switch bill {
        case 5:
            count5++
        case 10:
            count5--
            count10++
        case 20:
            if count10 > 0 {
                count10--
                count5--
            } else {
                count5 -= 3
            }
        default:
            return false
        }
        if count5 < 0 {
            return false
        }
    }
    return true
}
