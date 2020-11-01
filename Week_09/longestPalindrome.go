package Week_09

//时间复杂度O(n^2)，空间复杂度O(n), n为字符串长度
func longestPalindrome(s string) string {
    var cand string
    for i := 0; i < len(s); i++ {
        cand = checkPalindrome(s, i, i, cand)
        cand = checkPalindrome(s, i, i+1, cand)
    }
    return cand
}

func checkPalindrome(s string, left, right int, cand string) string {
    for {
        if !(0 <= left && right < len(s) && s[left] == s[right]) {
            break
        }
        left--
        right++
    }
    if len(s[left+1:right]) > len(cand) {
        cand = s[left+1:right]
    }
    return cand
}
