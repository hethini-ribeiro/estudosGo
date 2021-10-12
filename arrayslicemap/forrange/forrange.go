package main

import "fmt"

func main() {
	letras := [...]string{"a", "b", "c", "d", "e"} //alocação dinamica

	// range letras retorna dois valores do array letras: o indice e o conteudo
	for i, letra := range letras {
		fmt.Println("indice:", i)
		fmt.Println("Conteudo:", letra)
		fmt.Printf("%d) %s\n", i, letra)
	}

	// nesse FOR eu to eliminando a necessidade de pegar o indice
	for _, letra := range letras {
		fmt.Println("Conteudo:", letra)
	}
}
