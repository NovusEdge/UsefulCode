package sorting

import (
	"container/heap"
)

//HeapSort sorts an array using the ***heap sort algorithm***
func HeapSort(arr *[]int64) {
	if !(len(*arr) < 2) {

		res := []int64{}
		arrHeap := _makeHeap(*arr)

		for arrHeap.Len() > 0 {
			data, _ := heap.Pop(arrHeap).(int64)
			res = append(res, data)
		}

		*arr = res
	}
}

/*
HeapSorted sorts an array using the ***heap sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func HeapSorted(arr []int64) (res []int64) {
	if len(arr) < 2 {
		return arr
	}

	arrHeap := _makeHeap(arr)

	for arrHeap.Len() > 0 {
		data, _ := heap.Pop(arrHeap).(int64)
		res = append(res, data)
	}

	return
}

func _makeHeap(arr []int64) *intHeap {
	h := &intHeap{}
	heap.Init(h)
	for _, i := range arr {
		heap.Push(h, i)
	}
	return h
}

type intHeap []int64

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
