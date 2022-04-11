package main

import (
	"fmt"
)

// Struct
type Medico struct {
	numMed  int
	nomeMed string
	comp    []string
}

func main() {
	umMedico := Medico{
		numMed:  8,
		nomeMed: "João",
		comp: []string{
			"Maria",
			"Paulo",
		},
	}
	fmt.Println(umMedico)
}
