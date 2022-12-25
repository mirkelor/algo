package common

type Heap[T any] struct {
	Heap       []T
	Comparator func(i, j T) bool
}

func NewHeap[T any](comparator func(i, j T) bool) *Heap[T] {
	return &Heap[T]{Heap: make([]T, 1), Comparator: comparator}
}

func (h *Heap[T]) Len() int {
	return len(h.Heap) - 1
}

func (h *Heap[T]) Push(val T) {
	h.Heap = append(h.Heap, val)
	i := h.Len()
	// Percolate up
	percolateUp(h, i)
}

func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		return *new(T)
	}
	if h.Len() == 1 {
		res := h.Heap[1]
		h.Heap = h.Heap[:1]
		return res
	}

	res := h.Heap[1]
	// Move last value to root
	h.Heap[1], h.Heap[h.Len()] = h.Heap[h.Len()], h.Heap[1]
	h.Heap = h.Heap[:h.Len()]
	i := 1
	// Percolate down
	percolateDown(h, i)
	return res
}

func (h *Heap[T]) Peek() T {
	if h.Len() > 0 {
		return h.Heap[1]
	}
	return *new(T)
}

func (h *Heap[T]) Heapify(arr []T) {
	// 0-th position is moved to the end
	arr = append(arr, arr[0])
	arr[0] = *new(T)

	h.Heap = arr
	cur := h.Len() - 1

	for cur > 0 {
		percolateDown(h, cur)
		cur--
	}
}

func percolateUp[T any](h *Heap[T], i int) {
	for i/2 > 0 && h.Comparator(h.Heap[i], h.Heap[i/2]) {
		h.Heap[i], h.Heap[i/2] = h.Heap[i/2], h.Heap[i]
		i /= 2
	}
}

func percolateDown[T any](h *Heap[T], i int) {
	for i > 0 && 2*i <= h.Len() {
		l, r := 2*i, 2*i+1
		if r <= h.Len() && h.Comparator(h.Heap[r], h.Heap[l]) && h.Comparator(h.Heap[r], h.Heap[i]) {
			// Swap right child
			h.Heap[i], h.Heap[r] = h.Heap[r], h.Heap[i]
			i = r
		} else if h.Comparator(h.Heap[l], h.Heap[i]) {
			// Swap left child
			h.Heap[i], h.Heap[l] = h.Heap[l], h.Heap[i]
			i = l
		} else {
			break
		}
	}
}
