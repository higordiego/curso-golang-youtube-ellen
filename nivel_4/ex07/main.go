package main

import "fmt"

func main() {
	slice := [][]string{
		[]string{
			"Higor",
			"Diego",
			"poker",
		},
		[]string{
			"Diego",
			"Alves",
			"poker",
		},
		[]string{
			"Ferreira",
			"Alves",
			"poker",
		},
	}

	for _, s := range slice {
		for _, v := range s {
			fmt.Println("v", v)
			
		}
	}
}
