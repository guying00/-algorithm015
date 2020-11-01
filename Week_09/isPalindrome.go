package Week_09

//时间复杂度O(n)，空间复杂度O(n), n为字符串长度
func isPalindrome(s string) bool {
    str := preDeal(s)
    n := len(str)
    i, j := n/2, n/2
    if n % 2 == 0 {i--}
    for ;i >= 0; i, j = i-1, j+1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true
}

func preDeal(s string) string {
    r := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        switch {
        case '0' <= s[i] && s[i] <= '9':
            r = append(r, s[i])
        case 'a' <= s[i] && s[i] <= 'z':
            r = append(r, s[i])
        case 'A' <= s[i] && s[i] <= 'Z':
            r = append(r, s[i]+32)
        }
    }
    return string(r)
}
