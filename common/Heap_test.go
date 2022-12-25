package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	minHeap := NewHeap(func(i, j int) bool { return i < j })
	minHeap.Push(3)
	assert.Equal(t, 3, minHeap.Peek())

	minHeap.Push(1)
	assert.Equal(t, 1, minHeap.Peek())

	minHeap.Push(10)
	assert.Equal(t, 1, minHeap.Peek())
	assert.Equal(t, 3, minHeap.Len())
	assert.Equal(t, 1, minHeap.Pop())
	assert.Equal(t, 2, minHeap.Len())
}

func TestHeapify(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      func(i, j int) bool
		expected []int
	}{
		{[]int{1, 2, 3}, func(i, j int) bool { return i < j }, []int{0, 1, 3, 2}},
		{[]int{1, 2, 3}, func(i, j int) bool { return i > j }, []int{0, 3, 2, 1}},
		{[]int{60, 50, 80, 40, 30, 10, 70, 20, 90}, func(i, j int) bool { return i < j }, []int{0, 10, 30, 20, 50, 80, 70, 40, 90, 60}},
	}

	for _, tt := range tests {
		heap := NewHeap(tt.in2)
		heap.Heapify(tt.in1)
		assert.Equal(t, tt.expected, heap.Heap)
	}
}
