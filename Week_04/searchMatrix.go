package Week_04

//采用二分查找
//时间复杂度O(log(m*n))，空间复杂度O(1)
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }

    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, m * n
    mid := (left+right) / 2
    for left <= right && mid < m * n {
        i := mid / n
        j := mid % n
        if matrix[i][j] == target {
            return true
        }

        if matrix[i][j] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
        mid = (left+right) / 2
    }
    return false
}
