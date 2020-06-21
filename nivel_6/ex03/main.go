package main

import "fmt"

func deferTest() {
	defer fmt.Println("final")
	fmt.Println("Inicial")
}

func main() {
	deferTest()
}