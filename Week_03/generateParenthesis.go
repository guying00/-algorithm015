package Week_03

//时间复杂度：O(\dfrac{4^n}{\sqrt{n}})
//空间复杂度：O(n)
func generateParenthesis(n int) []string {
    result := []string{}
    cand := make([]byte, 2 * n)
    var helper func(curIdx, left, right int)
    helper = func(curIdx, left, right int) {
        if curIdx == 2*n {
            result = append(result, string(cand))
            return
        }

        if left < n {
            cand[curIdx] = '('
            helper(curIdx+1, left+1, right)
        }

        if right < left {
            cand[curIdx] = ')'
            helper(curIdx+1, left, right+1)
        }
    }
    helper(0, 0, 0)
    return result
}
