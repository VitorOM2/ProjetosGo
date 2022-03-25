package main

import (
	"fmt"
)

func main() {
	// ---------- Primeiro Exemplo ----------
	a := []int{1, 2, 3}
	b := a         // Cria um segundo slice que está apontando para o primeiro
	b[1] = 5       // O segundo slice modifica o segundo indice dos slices
	fmt.Println(a) // Slice modificado por causa do segundo slice
	fmt.Println(b)
	fmt.Printf("Tamanho:	%v \n", len(a))    // Mostra o tamanho do slice
	fmt.Printf("Capacidade:	%v \n", cap(a)) // Mostra a capacidade do slice

	fmt.Println("")

	// ---------- Segundo Exemplo ----------
	// Outras formas de criar slices

	sa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sb := sa[:]              // Cria um slice com todos elementos
	sc := sa[3:]             // Cria um slice do 4 elemento até o fim
	sd := sa[:6]             // Cria um slice com os primeiros 6 elementos
	se := sa[3:6]            // Cria um slice com o quarto, quinto e sexto elemento
	sf := make([]int, 3, 10) // Cria um slice usando o método make com o tamanho 3 e capacidade 10

	fmt.Printf("Slice A: %v \n", sa)
	fmt.Printf("Slice B: %v \n", sb)
	fmt.Printf("Slice C: %v \n", sc)
	fmt.Printf("Slice D: %v \n", sd)
	fmt.Printf("Slice E: %v \n", se)
	fmt.Printf("Slice F: %v \n", sf)
}
