package arrays

import (
	"container/heap"
)

func SecondHigeshtNumber(input []int) int {

	lar := 0
	sec := 0

	for _, val := range input {

		if val > lar { // 4 1 6 7 9
			sec = lar
			lar = val
		} else if val > sec && val < lar {
			sec = val
		}

	}
	return sec
}

type Heap []int

func (h Heap) Len() int {

	return len(h)

}

func (h Heap) Less(i int, j int) bool {

	return h[i] < h[j]

}

func (h Heap) Swap(i int, j int) {

	h[i], h[j] = h[j], h[i]

}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKHigeshtNumber(input []int, k int) int {
	h := Heap(input[:k])

	heap.Init(&h)

	for _, num := range input[k:] {

		if h[0] < num {
			heap.Pop(&h)
			heap.Push(&h, num)
		}

	}

	return h[0]

}
