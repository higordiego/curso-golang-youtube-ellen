package main

import "fmt"

func somaVariatica(numeros ...int) int {
	valor := 0
	for _, numero := range numeros {
		valor += numero
	}
	return valor
}

func somaVariaticaSlice(numeros []int) int {
	valor := 0
	for _, numero := range numeros {
		valor += numero
	}
	return valor
}

func main() {

	si := []int{10, 2, 30, 4, 20}

	fmt.Println(somaVariatica(si...))
	fmt.Println(somaVariaticaSlice(si))

}
