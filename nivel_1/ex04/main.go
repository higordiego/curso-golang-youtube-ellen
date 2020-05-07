package main

import (
	"fmt"
)

type test int

var x test

func main() {
	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)
}
