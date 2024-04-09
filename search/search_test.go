package search

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
	}

	for _, tt := range data {
		result := binarySearch(tt.arr, tt.target)
		if result != tt.expected {
			t.Errorf("binarySearch(%v, %d) = %d; expected %d", tt.arr, tt.target, result, tt.expected)
		}
	}
}
