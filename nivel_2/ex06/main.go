package main

import "fmt"

const (
	_ = 1994 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)

	fmt.Printf("%d", 31337)
}