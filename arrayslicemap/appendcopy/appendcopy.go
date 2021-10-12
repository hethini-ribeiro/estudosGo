package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) // append só trabalha com slice enão com array
	slice1 = append(slice1, 4, 5, 6) //add elementos dentro de um slice e se precisar, dobra o tamanho da capacidade
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) //cria lice de 2 de tamanho
	copy(slice2, slice1)     //copio 2 elementos de slice1 sem expandir a capacidade do slice
	fmt.Println(slice2)
}
