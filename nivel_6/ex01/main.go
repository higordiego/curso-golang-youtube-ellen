package main

import "fmt"

func retornaInt() int {
	return 1
}

func retornaIntString() (int, string) {
	return 2, "testando"
}

func main() {

	intA := retornaInt()

	inteiro, stringA := retornaIntString()

	fmt.Println(intA)
	fmt.Println(inteiro, stringA)
}
