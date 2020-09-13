package Week_01

func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }

    n := len(height)
    maxLeft := make([]int, n)
    maxRight := make([]int, n)
    maxLeft[0] = height[0]
    for i := 1; i < n; i++ {
        maxLeft[i] = max(maxLeft[i-1], height[i])
    }

    maxRight[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], height[i])
    }

    result := 0
    for i := 1; i <= n - 2; i++ {
        result += (min(maxLeft[i], maxRight[i]) - height[i])
    }
    return result
}

func trap1(height []int) int {
    result := 0
    for i := 1; i < len(height) - 1; i++ {
        maxLeft, maxRight := 0, 0
        for l := i-1; l >= 0;l-- {
            maxLeft = max(maxLeft, height[l])
        }
        for r := i+1; r < len(height); r++ {
            maxRight = max(maxRight, height[r])
        }

        if min(maxLeft, maxRight) > height[i] {
            result += (min(maxLeft, maxRight) - height[i])
        }
    }
    return result
}

func trap2(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }

    result := 0
    left, right := 0, n-1
    leftMax, rightMax := 0, 0
    for left < right {
        if height[left] < height[right] {
            if leftMax < height[left] {
                leftMax = height[left]
            } else {
                result += (leftMax - height[left])
            }
            left++
        } else {
            if rightMax < height[right] {
                rightMax = height[right]
            } else {
                result += (rightMax - height[right])
            }
            right--
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


