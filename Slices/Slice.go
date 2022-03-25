package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Tamanho:	%v \n", len(a)) // Mostra o tamanho do slice
	fmt.Printf("Capacidade:	%v", cap(a)) // Mostra a capacidade do slice
}
