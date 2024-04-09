package search

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := (low + high) / 2
		value := arr[mid]

		if value == target {
			return mid
		} else if value > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

// arr should be sorted
func binarySearchRecursive(arr []int, target int) int {
	if len(arr) <= 0 {
		return -1
	}

	midPoint := len(arr) / 2

	if midPoint <= 1 {
		return -1
	}

	if arr[midPoint] == target {
		return midPoint
	}

	if arr[midPoint] > target {
		binarySearch(arr[midPoint:], target)
	} else {
		binarySearch(arr[:midPoint+1], target)
	}

	return -1
}
