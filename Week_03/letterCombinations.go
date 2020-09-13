package Week_03

//递归回溯实现
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }

    result := []string{}
    cand := make([]byte, len(digits))
    var helper func(curIdx int)
    helper = func(curIdx int) {
        if curIdx == len(digits) {
            result = append(result, string(cand))
            return
        }

        for _, c := range numToBytes(digits[curIdx]) {
            cand[curIdx] = c
            helper(curIdx+1)
        }
    }
    helper(0)
    return result
}

func numToBytes(n byte) []byte {
    switch n {
    case '2':
        return []byte{'a', 'b', 'c'}
    case '3':
        return []byte{'d', 'e', 'f'}
    case '4':
        return []byte{'g', 'h', 'i'}
    case '5':
        return []byte{'j', 'k', 'l'}
    case '6':
        return []byte{'m', 'n', 'o'}
    case '7':
        return []byte{'p', 'q', 'r', 's'}
    case '8':
        return []byte{'t', 'u', 'v'}
    case '9':
        return []byte{'w', 'x', 'y', 'z'}
    }
    return nil
}

//迭代回溯实现
func letterCombinations1(digits string) []string {
    if digits == "" {
        return nil
    }
    r := &[]string{""}
    for _, d := range digits {
        r = genStr(r, d)
    }
    return *r
}

func genStr(curStr *[]string, c rune) *[]string {
    r := []string{}
    for _, s := range *curStr {
        for _, a := range numToChars(c) {
            r = append(r, s + string(a))
        }
    }
    return &r
}

func numToChars(c rune) string {
    switch c {
    case '2':
        return "abc"
    case '3':
        return "def"
    case '4':
        return "ghi"
    case '5':
        return "jkl"
    case '6':
        return "mno"
    case '7':
        return "pqrs"
    case '8':
        return "tuv"
    case '9':
        return "wxyz"
    }
    return ""
}