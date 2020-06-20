package main

import "fmt"

func main() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Higor",
		sabor:   "sovete",
		ondetem: []string{"centro", "bairro"},
		vaibemcom: map[string]string{
			"bebida": "acool",
		},
	}

	fmt.Println(x)

}
