package main

import (
	"fmt"
	"log"
)

type test int
var x test
var y int


func main() {
	log.Println(x)
	fmt.Printf("%T", x)
	y = int(x)
	log.Println(y)
	log.Printf("%T", y)
}