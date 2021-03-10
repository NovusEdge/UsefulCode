package main

import (
	"fmt"
	"sorting/sorting"
)

func main() {
	fmt.Println("\nSorting with Bubble Sort")
	unsorted := []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.BubbleSorted(unsorted))
	sorting.BubbleSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Insertion Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.InsertionSorted(unsorted))
	sorting.InsertionSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Count Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.CountSorted(unsorted))
	sorting.CountSort(&unsorted) // this one needs work ...
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Selection Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.SelectionSorted(unsorted))
	sorting.SelectionSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Heap Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.HeapSorted(unsorted))
	sorting.HeapSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Radix Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println("Using the 'Sorted' Method: ", sorting.RadixSorted(unsorted))
	sorting.RadixSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)

	fmt.Println("\nSorting with Merge Sort")
	unsorted = []int64{1, 4, 6, 43, 2, 4, 57, 8, 53, 1, 10, 123, 20, 24, 3}
	fmt.Println("Original: ", unsorted)
	// fmt.Println("Using the 'Sorted' Method: ", sorting.MergeSorted(unsorted))
	sorting.MergeSort(&unsorted)
	fmt.Println("Using the 'Sort' Method:   ", unsorted)
}
