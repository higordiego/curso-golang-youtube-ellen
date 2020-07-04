package main

import "golang.org/x/exp/errors/fmt"


type pessoa struct {
	FirstName string
	LastName string
}

type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Printf("Aqui quem  está falando é: %v %v", p.FirstName, p.FirstName)
}

func dizerAlgumaCoisa (h humanos) {
	h.falar()
}

func main() {
	pe := pessoa{"Higor", "Diego"}
	dizerAlgumaCoisa(&pe)
}
