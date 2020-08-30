package Week_01

type MinStack struct {
    data []int
    min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{}, []int{math.MaxInt32}}
}


func (this *MinStack) Push(x int) {
    this.data = append(this.data, x)
    if this.min[len(this.min)-1] >= x {
        this.min = append(this.min, x)
    }
}


func (this *MinStack) Pop() {
    if len(this.data) <= 0 {
        return
    }
    top := this.data[len(this.data)-1]
    this.data = this.data[:len(this.data)-1]
    if top == this.min[len(this.min)-1] {
        this.min = this.min[:len(this.min)-1]
    }
}


func (this *MinStack) Top() int {
    if len(this.data) <= 0 {
        return 0
    }
    return this.data[len(this.data)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.min) <= 1 {
        return 0
    }
    return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
