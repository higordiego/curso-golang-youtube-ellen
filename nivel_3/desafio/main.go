package main

import "fmt"

func main() {

	for i := 33; i < 123; i++ {
		fmt.Printf("%d %#x %#U \t \n ", i, i, i)
	}
}
