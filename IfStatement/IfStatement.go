package main

import (
	"fmt"
)

func main() {
	// Exemplo de If Statement com Boolean
	if true {
		fmt.Println("oi")
	}

	// Exemplo de If Statement com comparações

	num := 50
	palpite := 30

	if palpite < 1 || palpite > 10 {
		fmt.Println("O palpite tem que ser entre 1 e 100")
	} else {
		if palpite < num {
			fmt.Println("Muito Baixo")
		}
		if palpite > num {
			fmt.Println("Muito Alto")
		}
		if palpite == num {
			fmt.Println("Acertou")
		}
	}

}
