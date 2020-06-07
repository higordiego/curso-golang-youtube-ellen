package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// slice = append(slice, slice[2])
	// fmt.Println(slice)

	// slice = append(slice, slice[len(slice)-1])
	// fmt.Println(slice)

	// newSlice := slice[3:len(slice)-1]
	// slice = append(slice, newSlice[len(newSlice)-2])

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:8])
	fmt.Println(slice[2:len(slice)-1])
}
