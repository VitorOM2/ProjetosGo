package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tipos de dados: Boolean, tipos númericos e tipos de textos")

	// ---------- BOOLEAN ----------
	fmt.Println("Boolean: Verdadeiro ou falso")

	var bP bool = true
	var bN bool = false

	fmt.Printf("Exemplo de um boolean positivo: %v, %T \n", bP, bP)
	fmt.Printf("Exemplo de um boolean falso: %v, %T \n", bN, bN)
}
