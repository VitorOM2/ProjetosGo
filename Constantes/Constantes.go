package main

import (
	"fmt"
)

const (
	// Iota auto incrementa constantes
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

func main() {
	// const minha_const int = 8 // Exemplo de declaração de constante
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
}
