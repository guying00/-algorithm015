package Week_09

//时间复杂度O(n^2)，空间复杂度O(n), n为字符串长度
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    s2t := map[byte]byte{}
    t2s := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        c1, ok1 := s2t[s[i]]
        c2, ok2 := t2s[t[i]]
        if !ok1 && !ok2 {
            s2t[s[i]] = t[i]
            t2s[t[i]] = s[i]
        } else {
            if c1 != t[i] || c2 != s[i] {
                return false
            }
        }
    }
    return true
}
