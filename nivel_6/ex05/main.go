package main

import (
	"fmt"
	"math"
)

func main() {
	x := quadrado{15.0}
	y := ciculo{5.25}

	medida(x)
	medida(y)
}

type quadrado struct {
	lado float64
}
type ciculo struct {
	raio float64
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println(resultado)
}

func (c ciculo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println(resultado)
}
