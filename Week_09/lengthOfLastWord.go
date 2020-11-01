package Week_09

//时间复杂度O(n)，空间复杂度O(1), n为字符串长度
func lengthOfLastWord(s string) int {
    end := len(s)-1
    for ;end >= 0 && s[end] == ' '; end-- {}
    begin := end
    for ;begin >= 0 && s[begin] != ' '; begin-- {}
    return end - begin
}
