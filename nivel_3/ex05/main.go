package main

import "fmt"

func main() {

	for i := 10; i < 101; i++ {
		fmt.Printf("%v\t", i%4)
	}
}
