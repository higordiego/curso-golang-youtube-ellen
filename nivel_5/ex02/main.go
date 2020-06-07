package main

import "fmt"

type pessoa struct {
	Nome            string
	Sobrenome       string
	SaboresFavorits []string
}

func main() {

	mapPessoas := make(map[string]pessoa)
	mapPessoas["Juazeiro"] = pessoa{"Higor", "Diego", []string{"napolitano", "chocolate"}}
	mapPessoas["Barbalha"] = pessoa{"Alves", "Ferreira", []string{"napolitano", "chocolate"}}

	

	for key, v := range mapPessoas {
		fmt.Printf("cidade: %v \n", key)
		fmt.Printf("nome: %v \n", v.Nome)
		for _, s := range v.SaboresFavorits {
			fmt.Println("sabores: ", s)
		}
	}

}
