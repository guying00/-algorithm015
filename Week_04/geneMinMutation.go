package Week_04

func minMutation(start string, end string, bank []string) int {
    result := 0
    visited := make([]bool, len(bank))
    cache := []string{start}
    for len(cache) > 0 {
        curLen := len(cache)
        result++
        for _, curGene := range cache {
            for i := 0; i < len(bank); i++ {
                if visited[i] {
                    continue
                }
                if canTrans(curGene, bank[i]) {
                    if end == bank[i] {
                        return result
                    }
                    cache = append(cache, bank[i])
                    visited[i] = true
                }
            }
        }
        cache = cache[curLen:]
    }
    return -1
}

func canTrans(g1, g2 string) bool {
    if len(g1) != len(g2) {
        return false
    }
    count := 0
    for i := 0; i < len(g1); i++ {
        if g1[i] != g2[i] {
            if count > 0 {
                return false
            } else {
                count++
            }
        }
    }
    return count > 0
}
