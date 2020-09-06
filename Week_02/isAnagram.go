package Week_02

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    statis := map[rune]int{}
    for _, c := range s {
        statis[c]++
    }

    for _, c := range t {
        statis[c]--
        if statis[c] < 0 {
            return false
        }
    }
    return true
}