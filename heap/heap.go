package heap

type Heap struct {
	elements []int
}

func (h *Heap) Push(item int) {

	h.elements = append(h.elements, item)
	h.heapifyUp(len(h.elements) - 1)

}

func (h *Heap) Pop() int {

	max := h.elements[0]
	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]
	h.heapifyDown(0)

	return max

}

func (h *Heap) heapifyUp(index int) {

	for h.elements[parent(index)] < h.elements[index] {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

func (h *Heap) heapifyDown(index int) {

	lastIndex := len(h.elements) - 1
	l, r := left(index), right(index)
	childToComapre := 0
	for l <= lastIndex {

		if l == lastIndex {
			childToComapre = l
		} else if h.elements[l] > h.elements[index] {
			childToComapre = l
		} else {
			childToComapre = r
		}

		if h.elements[index] < h.elements[childToComapre] {
			h.Swap(index, childToComapre)
			index = childToComapre
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap) Swap(i int, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}
