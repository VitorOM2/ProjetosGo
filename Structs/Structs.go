package main

import (
	"fmt"
)

// Structs
// Exemplo 1
type Medico struct {
	numMed  int
	nomeMed string
	comp    []string
}

//Exemplo 2
type Animal struct {
	Nome   string
	Origem string
}

type Passaro struct {
	Animal
	VelocidadeKPH float32
	PodeVoar      bool
}

func main() {

	// ----- Exemplo 1 -----
	umMedico := Medico{
		numMed:  8,
		nomeMed: "João",
		comp: []string{
			"Maria",
			"Paulo",
		},
	}
	// ---------------------
	// ----- Exemplo 2 -----
	p := Passaro{}
	p.Nome = "Emu"
	p.Origem = "Austrália"
	p.VelocidadeKPH = 48
	p.PodeVoar = false

	fmt.Println(umMedico)
	fmt.Println(p)
}
