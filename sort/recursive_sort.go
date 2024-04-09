package sort

// ======================
// Mergesort
// ======================

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
}

func mergeSort(arr []int) []int {
	// base case
	if len(arr) <= 1 {
		return arr
	}

	// find the mid point
	midPoint := len(arr) / 2

	// divide the array
	left := mergeSort(arr[:midPoint])
	right := mergeSort(arr[midPoint:])

	return merge(left, right)
}

// ===============
// Quicksort
// ===============

func partition(arr []int, low, high int) int {
	// Pivot is the last element of the array
	pivot := arr[high]

	// This is what we will be returning
	index := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			// swap to the index position
			index++
			tmp := arr[index]
			arr[index] = arr[i]
			arr[i] = tmp
		}
	}

	index++
	arr[high] = arr[index]
	arr[index] = pivot

	return index
}

func qs(arr []int, low, high int) {

	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)

	// go on left and right side of the array
	// and exclude the pivot element
	qs(arr, low, pivotIndex-1)
	qs(arr, pivotIndex+1, high)
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
