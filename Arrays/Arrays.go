package main

import (
	"fmt"
)

func main() {
	// ----- Variáveis -----
	// Arrays
	var alunos [3]string
	notas := [...]int{8, 8, 10}

	// ----- Mostra no console -----
	fmt.Printf("Alunos %v \n", alunos)
	fmt.Printf("Notas: %v \n", notas)
	alunos[0] = "João"
	alunos[1] = "Maria"
	alunos[2] = "Ana"
	fmt.Printf("Alunos %v", alunos)
}
