package Week_01

func maxArea(height []int) int {
    if len(height) < 2 {
        return 0
    }

    curI, curJ := 0, len(height) - 1
    maxArea := -1
    maxLeft, maxRight := 1<<32 -1, 1<<32 -1
    for curI <= curJ {
        mid := (curJ - curI) * min(height[curI], height[curJ])
        if maxArea <= mid {
            maxArea = mid
            maxLeft = height[curI]
            maxRight = height[curJ]
        }

        if height[curI] < height[curJ] {
            curI++
            for curI < curJ && height[curI] <= maxLeft  {
                curI++
            }
        } else {
            curJ--
            for curI < curJ && height[curJ] <= maxRight  {
                curJ--
            }
        }
    }
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}