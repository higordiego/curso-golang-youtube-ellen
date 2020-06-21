package main


func somaVariatica (numeros ...int) int {
	valor := 0
	for _, numero := range numeros {
		valor += numero
	}
	return valor
}

func main() {
	
	
}