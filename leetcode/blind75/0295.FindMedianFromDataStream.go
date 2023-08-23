package blind75

import "github.com/mirkelor/algo/common"

type MedianFinder struct {
	MinHeap *common.Heap[int]
	MaxHeap *common.Heap[int]
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{
		MinHeap: common.NewHeap(func(i, j int) bool { return i < j }),
		MaxHeap: common.NewHeap(func(i, j int) bool { return i > j }),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.MaxHeap.Len() > 0 && num > mf.MaxHeap.Peek() {
		mf.MinHeap.Push(num)
	} else {
		mf.MaxHeap.Push(num)
	}

	for mf.MaxHeap.Len()-1 < mf.MinHeap.Len() {
		mf.MaxHeap.Push(mf.MinHeap.Pop())
	}

	for mf.MinHeap.Len() < mf.MaxHeap.Len()-1 {
		mf.MinHeap.Push(mf.MaxHeap.Pop())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.MaxHeap.Len() > mf.MinHeap.Len() {
		return float64(mf.MaxHeap.Peek())
	} else {
		return float64(mf.MaxHeap.Peek()+mf.MinHeap.Peek()) / 2.0
	}
}
