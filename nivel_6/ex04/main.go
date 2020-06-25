package main

import "fmt"

// Pessoa - struct
type Pessoa struct {
	Nome      string
	SobreNome string
	Idade     int
}

// NomeCompleto - metodo
func (p *Pessoa) NomeCompleto() {
	fmt.Printf("Nome: %v %v e sua Idade: %v", p.Nome, p.SobreNome, p.Idade)
}

func main() {

	tipoPessoa := Pessoa{"Higo", "Diego", 10}

	tipoPessoa.NomeCompleto()
}
