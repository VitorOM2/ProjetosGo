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
	fmt.Printf("Alunos %v \n", alunos) // Array sem nenhum item
	fmt.Printf("Notas: %v \n", notas)
	alunos[0] = "João" // Adiciona um item para a array
	alunos[1] = "Maria"
	alunos[2] = "Ana"
	fmt.Printf("Alunos %v \n", alunos)                  // Array atualizada com itens
	fmt.Printf("Quantidade de alunos: %v", len(alunos)) //Mostra a quantidade de itens em um array
}
