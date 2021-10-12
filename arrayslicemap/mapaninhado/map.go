package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 9564.56,
			"Guga Pereira":   4566.85,
		},
		"J": {
			"José João": 6566.84,
		},
		"P": {
			"Pedro Junior": 5948.56,
		},
	}

	delete(funcsPorLetra, "P") // deleta todos os valores dentro da chave "P"

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
