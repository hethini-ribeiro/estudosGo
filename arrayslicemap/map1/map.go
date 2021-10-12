package main

import "fmt"

func main() {
	// var aprovados map[int]string // Mapas não podem ser inicializado assim, pois ficam nulos
	aprovados := make(map[int]string) //Essa é a maneira correta de criar um map

	aprovados[12345678] = "Maria"
	aprovados[23456789] = "Pedro"
	aprovados[34567891] = "Ana"

	fmt.Println(aprovados)

	//range ta retornando indice e conteudo do map, no caso, cpf e nome
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println("Nome:", aprovados[12345678])
	delete(aprovados, 12345678)
	fmt.Println("Nome:", aprovados[12345678])
}
