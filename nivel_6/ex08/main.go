package main

import "fmt"

func main() {

	retornFuncao()()
}

func retornFuncao() func() {
	return func() {
		fmt.Println("testnado")
	}
}
