package main

import (
	"fmt"
)

func main() {
	sliceNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range sliceNumber {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", sliceNumber)
}
