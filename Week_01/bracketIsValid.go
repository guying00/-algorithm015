package Week_01

func isValid(s string) bool {
    if len(s) % 2 == 1 {
        return false
    }

    st := NewStack()
    for _, c := range s {
        top := st.Peek()
        if match[c] == 0 || top == 0 {
            st.Push(c)
            continue
        }

        if top != match[c] {
            return false
        } else {
            st.Pop()
        }
    }
    return st.Peek() == 0
}

var match = map[rune]rune{')': '(', '}': '{', ']':'['}

type stack struct {
    data []rune
    top int
}

func NewStack() *stack {
    return &stack{[]rune{}, -1}
}

func (s *stack) Push(d rune) {
    s.data = append(s.data, d)
    s.top++
}

func (s *stack) Pop() rune {
    if s.top < 0 {
        return 0
    }
    d := s.data[s.top]
    s.data = s.data[:s.top]
    s.top--
    return d
}

func (s *stack) Peek() rune {
    if s.top < 0 {
        return 0
    }
    return s.data[s.top]
}
