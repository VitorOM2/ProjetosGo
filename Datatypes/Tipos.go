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
	fmt.Printf("Exemplo de um boolean falso: %v, %T \n\n", bN, bN)

	// ---------- Tipos Numéricos ----------

	// ---------- Interger ----------
	fmt.Printf("Integer: Números inteiros													\n")
	fmt.Printf("Existem vários tipos de integer:											\n")
	fmt.Printf("Int8: pode ir do -128 até 127												\n")
	fmt.Printf("Int16: pode ir do -32.768 até 32.767										\n")
	fmt.Printf("Int32: pode ir do -2.147.483.648 até 2.147.483.647							\n")
	fmt.Printf("Int64: pode ir do -9.223.372.036.854.775.808 até 9.223.372.036.854.775.807	\n")
	fmt.Printf("Também existe os unsigned int:												\n")
	fmt.Printf("Uint8: pode ir do 0 até 255													\n")
	fmt.Printf("Uint16: pode ir do 0 até 65.535												\n")
	fmt.Printf("Uint32: pode ir do 0 até 4.294.967.295										\n")
}
