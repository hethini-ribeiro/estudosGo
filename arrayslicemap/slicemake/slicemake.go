package main

import "fmt"

func main() {
	s := make([]int, 10) //criando um slice de inteiros com 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //criando um slice de inteiros com 10 elementos e 20 posições
	// len é a qtd de elementos e cap é o tamanho real do slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) //anexo mais 10 elementos no fim do slice
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s)) //tem 20 elementos e 20 de capacidade

	s = append(s, 1)               //anexei mais um elemento no slice de capacidade 20
	fmt.Println(s, len(s), cap(s)) // ele add o elemento novo e dobra a capacidade do slice, indo de 20 pra 40
}
