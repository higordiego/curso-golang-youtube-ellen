package main

import "fmt"

func main() {
	x := 10
	if x > 12 && x < 20 {
		fmt.Printf("\t%v", x)
	} else {
		fmt.Printf("\t %v", x)
	}
}
