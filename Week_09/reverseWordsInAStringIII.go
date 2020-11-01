package Week_09

//时间复杂度O(n)，空间复杂度O(n), n为字符串长度
func reverseWords(s string) string {
    cache := make([]rune, len(s))
    count := 0
    for i, c := range s {
        if c != ' ' {
            count++
        } else {
            reverseOneWord(&cache, i-count, i-1)
            count = 0
        }
        cache[i] = c
    }
    reverseOneWord(&cache, len(s)-count, len(s)-1)
    return string(cache)
}

func reverseOneWord(s *[]rune, begin, end int) {
    for begin < end {
        m := (*s)[begin]
        (*s)[begin] = (*s)[end]
        (*s)[end] = m
        begin++
        end--
    }
}
