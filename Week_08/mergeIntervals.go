package Week_08

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    sort.Slice(intervals, func (i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    result := [][]int{}
    for _, val := range intervals {
        mergeFlag := false
        for idx, res := range result {
            if val[0] <= res[1] {
                result[idx][1] = max(val[1], res[1])
                mergeFlag = true
                break
            }
        }
        if !mergeFlag {
            result = append(result, val)
        }
    }
    return result
}

func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
