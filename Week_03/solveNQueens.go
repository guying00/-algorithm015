package Week_03

//时间复杂度：O(N!)
//空间复杂度：O(N)
func solveNQueens(n int) [][]string {
    result := [][]string{}
    queens := make([]int, n)

    columns := map[int]bool{}
    diagonals1 := map[int]bool{}
    diagonals2 := map[int]bool{}

    var helper func(row int)
    helper = func(row int) {
        if row == n {
            result = append(result, genScheme(queens))
            return
        }

        for col := 0; col < n; col++ {
            diag1 := row+col
            diag2 := row-col
            if columns[col] || diagonals1[diag1] || diagonals2[diag2] {
                continue
            }

            queens[row] = col
            columns[col] = true
            diagonals1[diag1] = true
            diagonals2[diag2] = true
            helper(row+1)
            columns[col] = false
            diagonals1[diag1] = false
            diagonals2[diag2] = false
        }
    }

    helper(0)
    return result
}

func genScheme(qs []int) []string {
    scheme := []string{}
    n := len(qs)
    for _, pos := range qs {
        r := make([]byte, n)
        for i := 0; i < n; i++ {
            r[i] = '.'
        }
        r[pos] = 'Q'
        scheme = append(scheme, string(r))
    }
    return scheme
}
