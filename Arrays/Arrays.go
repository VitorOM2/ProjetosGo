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

	// 1# Primeiro exemplo
	fmt.Printf("Alunos %v \n", alunos) // Array sem nenhum item
	fmt.Printf("Notas: %v \n", notas)
	alunos[0] = "João" // Adiciona um item para a array
	alunos[1] = "Maria"
	alunos[2] = "Ana"
	fmt.Printf("Alunos %v \n", alunos)                    // Array atualizada com itens
	fmt.Printf("Quantidade de alunos: %v\n", len(alunos)) //Mostra a quantidade de itens em um array

	// 2# Segundo exemplo
	var superArray [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(superArray)
}
