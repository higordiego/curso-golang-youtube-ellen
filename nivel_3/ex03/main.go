package main

import "fmt"

func main() {
	ano_nacimento := 1987
	ano_atual := 2021

	for ano_nacimento < ano_atual {

		fmt.Printf("%v ", ano_nacimento)
		ano_nacimento++
	}
}
