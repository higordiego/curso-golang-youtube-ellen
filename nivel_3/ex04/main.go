package main

import "fmt"

func main() {
	ano_nacimento := 1987
	ano_atual := 2021

	for {
		if ano_nacimento < ano_atual {
			fmt.Println(ano_nacimento)
			ano_nacimento++
		}
		if ano_nacimento == ano_atual {
			break
		}
	}
}
