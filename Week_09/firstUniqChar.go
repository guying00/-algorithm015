package Week_09

//时间复杂度O(len(s))，空间复杂度O(1)
func firstUniqChar(s string) int {
    statis := make([]int, 26)
    for i := 0; i < len(s); i++ {
        statis[s[i]-'a']++
    }

    for i := 0; i < len(s); i++ {
        if statis[s[i]-'a'] == 1 {
            return i
        }
    }
    return -1
}

//时间复杂度O(len(s)^2)，空间复杂度O(1)
func firstUniqChar1(s string) int {
    for i := 0; i < len(s); i++ {
        j := 0
        for ; j < len(s); j++ {
            if i != j && s[i] == s[j] {
                break
            }
        }
        if j == len(s) {
            return i
        }
    }
    return -1
}
