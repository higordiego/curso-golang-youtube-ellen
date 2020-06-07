package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	v1 := caminhonete{veiculo{4, "azul"}, false}

	v2 := sedan{veiculo{4, "vermelha"}, true}

	fmt.Println(v1, v2)
}
