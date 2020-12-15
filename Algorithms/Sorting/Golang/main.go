package main

import (
	"fmt"
	"sorting/sorting"
)

func main() {
	fmt.Println("\nSorting with Bubble Sort")
	unsorted := []int{1, 4, 6, 43, 2, 4, 57, 8, 53, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println(sorting.BubbleSorted(unsorted))
	sorting.BubbleSort(&unsorted)
	fmt.Println(unsorted)

	fmt.Println("\nSorting with Insertion Sort")
	unsorted = []int{1, 4, 6, 43, 2, 4, 57, 8, 53, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println(sorting.InsertionSorted(unsorted))
	sorting.InsertionSort(&unsorted)
	fmt.Println(unsorted)

	fmt.Println("\nSorting with Count Sort")
	unsorted = []int{1, 4, 6, 43, 2, 4, 57, 8, 53, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println(sorting.CountSorted(unsorted))
	sorting.CountSort(&unsorted)
	fmt.Println(unsorted)

	fmt.Println("\nSorting with Selection Sort")
	unsorted = []int{1, 4, 6, 43, 2, 4, 57, 8, 53, 24, 3}
	fmt.Println("Original: ", unsorted)
	fmt.Println(sorting.SelectionSorted(unsorted))
	sorting.SelectionSort(&unsorted)
	fmt.Println(unsorted)
}
