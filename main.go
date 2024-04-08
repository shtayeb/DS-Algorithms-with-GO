package main

import "fmt"

func main() {
	println("this is test")

	arr := []int{1, 2, 3, 4, 5}

	mid := len(arr) / 2

	fmt.Printf("%v", arr[:mid+1])

	fmt.Printf("%v", arr[mid+1:])
}
