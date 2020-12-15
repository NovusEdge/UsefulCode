package sorting

//InsertionSort sorts an array using the ***insertion sort algorithm***
func InsertionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr); j++ {
			if (*arr)[j] > (*arr)[i] {
				(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			}
		}
	}
}

/*
InsertionSorted sorts an array using the ***insertion sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func InsertionSorted(arr []int) []int {
	sortedArr := arr
	for i := 0; i < len(sortedArr); i++ {
		for j := 0; j < len(sortedArr); j++ {
			if sortedArr[j] > sortedArr[i] {
				sortedArr[j], sortedArr[i] = sortedArr[i], sortedArr[j]
			}
		}
	}
	return sortedArr
}
