package sort

func bubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				// swap
				tmp := input[j]
				input[j] = input[j+1]
				input[j+1] = tmp
			}
		}
	}
}

func insertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		for j := i; j > 0 && input[j-1] > input[j]; j-- {
			input[j], input[j-1] = input[j-1], input[j]
		}
	}
}
