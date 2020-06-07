package main

import (
	"fmt"
	"log"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T \n", array)
	for i, v := range array {
		log.Println(i, v)
	}
}
