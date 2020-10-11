package Week_07

func ladderLength(beginWord string, endWord string, wordList []string) int {
    countB := 0
    countE := 0
    visitedB := make([]bool, len(wordList))
    visitedE := make([]bool, len(wordList))
    exist := false
    //先判断结束此是否存在字典中，不存在直接返回0, 存在则将从结尾搜索的访问数组对应位置true, 避免当开始词可以直接转换到结束词时, 计数错误(会多1次计数)
    for i, word := range wordList {
        if word == endWord {
            exist = true
            visitedE[i] = true
            break
        }
    }
    if !exist {
        return 0
    }

    cacheB := []*string{&beginWord}
    cacheE := []*string{&endWord}
    helper := func(cache *[]*string, level []*string, visited *[]bool, count *int) bool {
        *count++
        for idx := 0; idx < len(level); idx++ {
            curWord := level[idx]
            for i := 0; i < len(wordList); i++ {
                if (*visited)[i] {
                    continue
                }

                word := &wordList[i]
                if canTrans(curWord, word) {
                    (*visited)[i] = true
                    *cache = append(*cache, word)

                    //两个访问数组对应为都已经置true, 说明两边的搜索已经相遇, 返回true结束搜索
                    if visitedB[i] && visitedE[i] {
                        *count++
                        return true
                    }
                }
            }
        }
        return false
    }

    for len(cacheB) > 0 && len(cacheE) > 0 {
        //从开始词搜索
        curLen := len(cacheB)
        if helper(&cacheB, cacheB[:curLen], &visitedB, &countB) {
            return countB+countE
        }
        cacheB = cacheB[curLen:]

        //从结束词搜索
        curLen = len(cacheE)
        if helper(&cacheE, cacheE[:curLen], &visitedE, &countE) {
            return countB+countE
        }
        cacheE = cacheE[curLen:]
    }

    return 0
}

//判断两个单词是否能够转换
func canTrans(b, e *string) bool {
    if len(*b) != len(*e) {
        return false
    }
    count := 0
    for i := 0; i < len(*b); i++ {
        if (*b)[i] != (*e)[i] {
            count++
        }
        if count > 1 {
            return false
        }
    }
    return true
}
