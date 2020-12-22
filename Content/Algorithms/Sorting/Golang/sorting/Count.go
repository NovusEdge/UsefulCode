package sorting

//CountSort sorts an array using the ***count sort algorithm***
func CountSort(arr *[]int64) {
	resArr := []int64{}
	maxElem := _max(*arr)
	counter := make([]int64, maxElem+1)
	for _, i := range *arr {
		counter[i]++
	}
	for i := int64(0); i < int64(len(counter)); i++ {
		for j := int64(0); j < counter[i]; j++ {
			resArr = append(resArr, i)
		}
	}
	*arr = resArr
}

/*
CountSorted sorts an array using the ***count sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func CountSorted(arr []int64) []int64 {
	resArr := []int64{}
	maxElem := _max(arr)
	counter := make([]int64, maxElem+1)
	for _, i := range arr {
		counter[i]++
	}
	for i := int64(0); i < int64(len(counter)); i++ {
		for j := int64(0); j < counter[i]; j++ {
			resArr = append(resArr, i)
		}
	}
	return resArr
}

func _max(arr []int64) int64 {
	maxVal := arr[0]
	for _, i := range arr {
		if i > maxVal {
			maxVal = i
		}
	}
	return maxVal
}
