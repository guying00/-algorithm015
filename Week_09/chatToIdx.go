package Week_09

var stateMap = [][]int{{0, 1, 2, 3}, {3, 3, 2, 3}, {3, 3, 2, 3}, {3, 3, 3, 3}}   //0: start, 1: signed+, 1: signed-, 2: innumber, 3: end

//时间复杂度O(n)，空间复杂度O(1), n为字符串长度
func chatToIdx(c byte) int {
    switch {
    case c == ' ':
        return 0
    case c == '+' || c == '-':
        return 1
    case '0' <= c && c <= '9':
        return 2
    }
    return 3
}

func myAtoi(str string) int {
    result := 0
    curState := 0 //start
    sign := 1
    for i := 0; i < len(str) && curState != 3; i++ {//3 is end
        c := str[i]
        curState = stateMap[curState][chatToIdx(c)]
        if curState == 2 {//innumber
            curNum := int(c - '0')
            if sign == 1 && (result > math.MaxInt32/10 || result == math.MaxInt32/10 && curNum > math.MaxInt32%10) {
                result = math.MaxInt32
                break
            } else if sign == -1 && (result > math.MaxInt32/10 || result == math.MaxInt32/10 && curNum > math.MaxInt32%10) {
                result = math.MinInt32
                sign = 1
                break
            }
            result = result * 10 + curNum
        } else if curState == 1 && c == '-' {//signed
            sign = -1
        }
    }
    return sign * result
}
