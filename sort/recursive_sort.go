package sort

func merge(left, right []int) []int {
	sorted := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sorted = append(sorted, left[0])
			// cut the slice
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			// cut the slice
			right = right[1:]
		}
	}

	return append(append(sorted, left...), right...)
	// return sorted
}

func mergeSort(arr []int) []int {
	// base case
	if len(arr) == 1 {
		return arr
	}

	// find the mid point
	midPoint := len(arr) / 2

	// divide the array
	left := mergeSort(arr[:midPoint])
	right := mergeSort(arr[midPoint:])

	return merge(left, right)
}
