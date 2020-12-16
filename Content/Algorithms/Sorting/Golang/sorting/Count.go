package sorting

//CountSort sorts an array using the ***count sort algorithm***
func CountSort(arr *[]int) {
	resArr := []int{}
	maxElem := _max(*arr)
	counter := make([]int, maxElem+1)
	for _, i := range *arr {
		counter[i]++
	}
	for i := 0; i < len(counter); i++ {
		for j := 0; j < counter[i]; j++ {
			resArr = append(resArr, i)
		}
	}
	*arr = resArr
}

/*
CountSorted sorts an array using the ***count sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func CountSorted(arr []int) []int {
	resArr := []int{}
	maxElem := _max(arr)
	counter := make([]int, maxElem+1)
	for _, i := range arr {
		counter[i]++
	}
	for i := 0; i < len(counter); i++ {
		for j := 0; j < counter[i]; j++ {
			resArr = append(resArr, i)
		}
	}
	return resArr
}

func _max(arr []int) int {
	maxVal := arr[0]
	for _, i := range arr {
		if i > maxVal {
			maxVal = i
		}
	}
	return maxVal
}
