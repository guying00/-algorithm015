package Week_09

//时间复杂度O(n)，空间复杂度O(1), n为字符串长度
func reverseString(s []byte)  {
    for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}
