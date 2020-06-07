package main

import "fmt"

type pessoa struct {
	Nome            string
	Sobrenome       string
	SaboresFavorits []string
}

func main() {
	pessoa1 := pessoa{"Higor", "Diego", []string{"napolitano", "chocolate"}}
	pessoa2 := pessoa{"Alves", "Ferreira", []string{"napolitano", "chocolate"}}

	slice := []pessoa{pessoa1, pessoa2}

	for _, v := range slice {
		fmt.Println(v.Nome)
		for _, s := range v.SaboresFavorits {
			fmt.Println("sabores: ", s)
		}
	}

}
