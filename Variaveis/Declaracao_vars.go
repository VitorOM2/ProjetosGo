package main

// ---------- IMPORTAÇÕES ----------
import (
	"fmt"
)

// ---------- FUNÇÕES ----------
func main() {

	// Formas de declarar variáveis:

	// 1º Exemplo:	VAR NomeVar Tipo
	//		NomeVar = Valor
	var x int
	x = 42
	fmt.Println("Exemplo de declarar variável nº 1. VAR NomeVar Tipo")
	fmt.Println(x)

	// 2º Exemplo:	Var NomeVar Tipo = Valor
	var y int = 8
	fmt.Println("\nExemplo de declarar variável nº 2. VAR NomeVar Tipo = Valor")
	fmt.Println(y)

	// 3º Exemplo: NomeVar := Valor
	z := 12
	fmt.Println("\nExemplo de declarar variável nº 3. NomeVar := Valor")
	fmt.Println(z)

	// Exemplos  de blocos de Variáveis

	// 1º Bloco
	var (
		nome      string = "José"
		sobrenome string = "Silva"
		id        int    = 1
	)

	// 2º Bloco

	var (
		contador int = 0
	)

	fmt.Println("\nPrimeiro Bloco:")
	fmt.Println(nome, sobrenome, id)

	fmt.Println("\nSegundo Bloco:")
	fmt.Println(contador)

	// Exemplo de converter funções
	fmt.Println("\nExemplo de converter funções: Int para Float: VarConvertida := Tipo(varParaSerConvertida) ")
	j := float32(x)
	fmt.Println(j)
}
