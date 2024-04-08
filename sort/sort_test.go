package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		// {[]int{5, 2, 8, 3, 1, 9, 4, 7, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		// {[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		// {[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 4, 5, 3, 2, 1}, []int{1, 2, 3, 4, 5, 9}},
		// {[]int{}, []int{}},
		// {[]int{1}, []int{1}},
	}
	for _, tt := range tests {
		res := mergeSort(tt.input)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("got = %v, expected %v", tt.input, tt.expected)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	t.Skip("skip bubble sort")
	var tests = []struct {
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

func TestInsertionSort(t *testing.T) {
	t.Skip("skip insertion sort")
	var tests = []struct {
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
		insertionSort(tt.input)
		if !reflect.DeepEqual(tt.input, tt.expected) {
			t.Errorf("got = %v, expected %v", tt.input, tt.expected)
		}
	}
}
