package Week_07

func solveNQueens(n int) [][]string {
    result := [][]string{}
    queens := make([]int, n)
    columns := map[int]bool{}
    diagonal1 := map[int]bool{}
    diagonal2 := map[int]bool{}
    var helper func(curRow int)
    helper = func(curRow int) {
        if curRow == n {
            result = append(result, genScheme(queens))
            return
        }

        for col := 0; col < n; col++ {
            diag1 := curRow+col
            diag2 := curRow-col
            if columns[col] || diagonal1[diag1] || diagonal2[diag2] {
                continue
            }

            queens[curRow] = col
            columns[col], diagonal1[diag1], diagonal2[diag2] = true, true, true
            helper(curRow+1)
            columns[col], diagonal1[diag1], diagonal2[diag2] = false, false, false
        }
    }
    helper(0)
    return result
}

func genScheme(queens []int) []string {
    scheme := make([]string, len(queens))
    row := make([]byte, len(queens))
    for i, pos := range queens {
        for i := 0; i < len(queens); i++ {
            row[i] = '.'
        }
        row[pos] = 'Q'
        scheme[i] = string(row)
    }
    return scheme
}
