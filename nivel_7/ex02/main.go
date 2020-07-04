package main

import "fmt"

type pessoa struct {
	endereco string
}

func (p *pessoa) mudaMe() {
	p.endereco = "testando"
}

func main() {
	pessoaX := pessoa{}

	pessoaX.mudaMe()
	fmt.Println(pessoaX.endereco)
}