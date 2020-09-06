package Week_02

import "container/heap"

func topKFrequent(nums []int, k int) []int {
    statis := map[int]int{}
    for _, n := range nums {
        statis[n]++
    }

    maxHeap := NewKMaxFreqHeap(k)
    heap.Init(maxHeap)
    for k, v := range statis {
        heap.Push(maxHeap, &frequency{k, v})
    }

    result := []int{}
    for maxHeap.Len() > 0 {
        f := heap.Pop(maxHeap)
        result = append(result, f.(*frequency).elem)
    }
    return result
}

type frequency struct {
    elem int
    count int
}

type kMaxFreqHeap struct {
    data []*frequency
    len  int
    cap  int
}

func NewKMaxFreqHeap(k int) *kMaxFreqHeap {
    return &kMaxFreqHeap{make([]*frequency, k), 0, k}
}

func (h *kMaxFreqHeap) Len() int {
    return h.len
}

func (h *kMaxFreqHeap) Less(i, j int) bool {
    return h.data[i].count < h.data[j].count
}

func (h *kMaxFreqHeap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *kMaxFreqHeap) Push(x interface{}) {
    if v, ok := x.(*frequency); ok {
        if ok && h.len >= h.cap && h.data[0].count < v.count {
            heap.Pop(h)
        }
        if h.len < h.cap {
            h.data[h.len] = v
            h.len++
        }
    }
}

func (h *kMaxFreqHeap) Pop() interface{} {
    if h.len > 0 {
        h.len--
        return h.data[h.len]
    }
    return nil
}
