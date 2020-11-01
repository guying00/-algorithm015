package Week_09

//时间复杂度O(n*m*log(m))，空间复杂度O(n*m), n为字符串列表长度, m为字符串平均长度
func groupAnagrams(strs []string) [][]string {
    cache := map[string][]string{}
    for _, str := range strs {
        s := []byte(str)
        sort.Slice(s, func(i, j int) bool {
            return s[i] < s[j]
        })
        cache[string(s)] = append(cache[string(s)], str)
    }

    result := [][]string{}
    for _, slice := range cache {
        result = append(result, slice)
    }
    return result
}
