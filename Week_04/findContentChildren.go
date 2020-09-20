package Week_04

//应用贪心法，从胃口最小的小孩开始，每次将最小的能满足当前孩子的饼干分发给他，然后继续处理剩下的小孩和饼干
//时间复杂度O(n)，空间复杂度O(1)
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    content := 0
    for gIdx, sIdx := 0, 0; gIdx < len(g) && sIdx < len(s); {
        if g[gIdx] > s[sIdx] {
            sIdx++
        } else {
            gIdx++
            sIdx++
            content++
        }
    }
    return content
}
