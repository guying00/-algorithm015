package Week_07

func solveSudoku(board [][]byte)  {
    n := len(board)
    rowCache := make([]map[byte]bool, n)
    colCache := make([]map[byte]bool, n)
    boxCache := make([]map[byte]bool, n)
    for i := 0; i < n; i++ {
        rowCache[i] = map[byte]bool{}
        colCache[i] = map[byte]bool{}
        boxCache[i] = map[byte]bool{}
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            val := board[i][j]
            if val != '.' {
                boxIdx := i/3 * 3 + j/3
                rowCache[i][val] = true
                colCache[j][val] = true
                boxCache[boxIdx][val] = true
            }
        }
    }

    var helper func(idx int) bool
    helper = func(idx int) bool {
        if idx == n*n {
            return true
        }

        i := idx/n
        j := idx%n
        if board[i][j] == '.' {
            boxIdx := i/3 * 3 + j/3
            for k := 0; k < 9; k++ {
                val := byte(k + '1')
                if rowCache[i][val] || colCache[j][val] || boxCache[boxIdx][val] {
                    continue
                }
                rowCache[i][val] = true
                colCache[j][val] = true
                boxCache[boxIdx][val] = true
                board[i][j] = val
                if helper(idx+1) {
                    return true
                }
                rowCache[i][val] = false
                colCache[j][val] = false
                boxCache[boxIdx][val] = false
                board[i][j] = '.'
            }
            return false
        }
        return helper(idx+1)
    }
    helper(0)
}
