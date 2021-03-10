package sorting

//RadixSort sorts an array using the ***radix sort algorithm***
func RadixSort(arr *[]int64) {

}

/*
RadixSorted sorts an array using the ***radix sort algorithm*** and returns the sorted array

! Note: This function doesn't **modify** the array passed into the function
*/
func RadixSorted(arr []int64) []int64 {
	res := arr
	maxElem := _max(res)

	place := int64(1)
	for int(maxElem/place) > 0 { //needs work ...
		_countingSort(res, place)
		place *= 10
	}

	return res
}

func _countingSort(arr []int64, place int64) []int64 {
	temp := arr
	size := len(temp)
	res := make([]int64, size)
	count := make([]int64, 10)

	for i := 0; i < size; i++ {
		index := int64(arr[i] / place)
		count[index%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	i := size - 1
	for i >= 0 {
		index := int64(arr[i] / place)
		res[count[index%10]-1] = arr[i]
		count[index%10]--
		i--
	}

	return arr
}
