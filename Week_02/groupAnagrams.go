package Week_02

import "sort"

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
