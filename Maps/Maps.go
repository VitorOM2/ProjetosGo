package main

import (
	"fmt"
)

func main() {
	// Exemplo de Map
	// Nome           // [tipo da key] tipos dos dados
	populacoesEstados := make(map[string]int)
	populacoesEstados = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(populacoesEstados)
}
