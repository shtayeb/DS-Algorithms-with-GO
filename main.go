package main

import (
	"fmt"

	"github.com/shtayeb/DS-Algorithms-with-GO/list"
)

func main() {
	println("This is test.")

	// Create a new linked list
	l := list.NewLinkedList()

	// Test appending elements
	l.Append(1)
	l.Append(2)
	l.Append(3)

	for i := l.Size - 1; i >= 0; i-- {
		fmt.Printf("i = %v \n", i)
	}
}
