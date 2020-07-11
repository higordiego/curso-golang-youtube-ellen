package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	b := make(chan int, 1)

	go func() {
		b <- 2
	}()

	fmt.Println(<-b)

}
