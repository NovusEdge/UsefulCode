package sorting

//SelectionSort sorts an array using the ***selection sort algorithm***
func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[i] {
				(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			}
		}
	}
}

/*
SelectionSorted sorts an array using the ***selection sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func SelectionSorted(arr []int) []int {
	resArr := arr
	for i := 0; i < len(resArr)-1; i++ {
		for j := i + 1; j < len(resArr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return resArr
}
