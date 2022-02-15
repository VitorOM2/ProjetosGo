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

}
