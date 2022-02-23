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

	fmt.Printf("Exemplo de um boolean positivo: %v, %T\n", bP, bP)
	fmt.Printf("Exemplo de um boolean falso: %v, %T \n\n", bN, bN)

	// ---------- TIPOS NUMÉRICOS ----------

	// ---------- Interger ----------
	fmt.Println("Integer: Números inteiros			")
	fmt.Println("Existem vários tipos de integer:	")
	fmt.Println("Int8: pode ir do -128 até 127		")
	fmt.Println("Int16: pode ir do -32.768 até 32.767										")
	fmt.Println("Int32: pode ir do -2.147.483.648 até 2.147.483.647							")
	fmt.Println("Int64: pode ir do -9.223.372.036.854.775.808 até 9.223.372.036.854.775.807	")
	fmt.Println("Também existe os unsigned int:		")
	fmt.Println("Uint8: pode ir do 0 até 255		")
	fmt.Println("Uint16: pode ir do 0 até 65.535	")
	fmt.Printf("Uint32: pode ir do 0 até 4.294.967.295										\n\n")

	// ---------- Float ----------
	fmt.Println("Float: Números decimais")
	fmt.Println("Existem alguns tipos de float")
	fmt.Println("Float32: vai de +-1.18E-38  até +-3.4E38 ")
	fmt.Println("Float64: vai de +-2.23E-308 até +-1.8E308")
	fmt.Printf("Exemplo de número float: 3.14\n\n")

	// --------- Complex ---------
	vC := 1 + 2i
	vC2 := 2 + 5.2i

	fmt.Println("O tipo de dado complex é a combinação de números reais e imaginários")
	fmt.Println("Existe Complex62 e Complex128")
	fmt.Println("Exemplo de números complexo: (1 + 2i) + (2 + 5.2i):", vC+vC2)

	// ---------- TIPOS TEXTUAIS ----------

	// --------- String ----------
	s := "Isso é uma string"

	fmt.Println("\nString: são feitas de bytes")
	fmt.Printf("%v, %T \n", s, s) // %v = Variável. %T = Tipo da variável

}
