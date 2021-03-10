package sorting

//BubbleSort sorts an array using the ***bubble sort algorithm***
func BubbleSort(arr *[]int64) {

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*(arr))[j] > (*(arr))[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

/*
BubbleSorted sorts an array using the ***bubble sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func BubbleSorted(arr []int64) []int64 {
	sortedArr := arr

	for i := 0; i < len(sortedArr); i++ {
		for j := 0; j < len(sortedArr)-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	return sortedArr
}
