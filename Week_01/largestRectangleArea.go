package Week_01

func largestRectangleArea(heights []int) int {
    maxArea := 0
    stack := NewStack()
    updateArea := func(rightBound int) int {
        rectInfo := stack.Pop()
        area := rectInfo.height * (rightBound - rectInfo.leftBound)
        if area > maxArea {
            maxArea = area
        }
        return rectInfo.leftBound
    }

    for curIdx, curHeight := range heights {
        leftBound := curIdx
        for !stack.IsEmpty() && stack.Peek().height > curHeight {
            leftBound = updateArea(curIdx)
        }
        stack.Push(&rectInfo{leftBound, curHeight})
    }

    for !stack.IsEmpty() {
        _ = updateArea(len(heights))
    }
    return maxArea
}

type rectInfo struct {
    leftBound int
    height int
}

type stack struct {
    data []*rectInfo
    top  int
}

func NewStack() *stack {
    return &stack{[]*rectInfo{}, -1}
}

func (s *stack) Push(d *rectInfo) {
    s.data = append(s.data, d)
    s.top++
}

func (s *stack) Pop() *rectInfo {
    if s.top < 0 {
        return nil
    }
    d := s.data[s.top]
    s.data = s.data[:s.top]
    s.top--
    return d
}

func (s *stack) Peek() *rectInfo {
    if s.top < 0 {
        return nil
    }
    return s.data[s.top]
}

func (s *stack) IsEmpty() bool {
    return s.top < 0
}
