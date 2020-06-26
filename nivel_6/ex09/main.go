package main

import "fmt"

func main() {
	x := func(a int64) {
		fmt.Println(a)
	}
	callback(x)
}

func callback(x func(a int64)) {
	x(1)
}
