package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := a         // Cria um segundo slice que está apontando para o primeiro
	b[1] = 5       // O segundo slice modifica o segundo indice dos slices
	fmt.Println(a) // Slice modificado por causa do segundo slice
	fmt.Println(b)
	fmt.Printf("Tamanho:	%v \n", len(a)) // Mostra o tamanho do slice
	fmt.Printf("Capacidade:	%v", cap(a)) // Mostra a capacidade do slice
}
