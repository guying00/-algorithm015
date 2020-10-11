package Week_07

func isValidSudoku(board [][]byte) bool {
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
            if val == '.' {
                continue
            }

            boxIdx := i/3 * 3 + j/3
            if rowCache[i][val] || colCache[j][val] || boxCache[boxIdx][val] {
                return false
            }
            rowCache[i][val] = true
            colCache[j][val] = true
            boxCache[boxIdx][val] = true
        }
    }
    return true
}
