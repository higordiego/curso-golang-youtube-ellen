package main

import "fmt"

func main() {

	mapNames := map[string][]string{
		"diego": {"poker", "jogar bola"},
		"alves": {"poker"},
	}


	mapNames["ferreira"] = []string{"poker", "jogar bola"}


	delete(mapNames, "diego")
	for k, val := range mapNames {
		fmt.Println(k)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}
}
