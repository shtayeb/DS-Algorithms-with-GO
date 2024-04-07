package sort

import (
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 8, 3, 1, 9, 4, 7, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, tt := range tests {
		bubbleSort(tt.input)
		if !reflect.DeepEqual(tt.input, tt.expected) {
			t.Errorf("got = %v, expected %v", tt.input, tt.expected)
		}
	}
}
