package main

import "fmt"

func main() {
	
	c := func() {
		fmt.Println("testando")
	}

	fmt.Println(retornaFuncao(c)())
}


func retornaFuncao(x func()) func() int {
	a := 30
	return func () int {
		return a
	}
}