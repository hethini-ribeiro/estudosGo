package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Slice é como um ponteiro para o elemento do array e o tamanho da "fatia"

	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	// Slice não é array! Slice deine um pedaço de um array
	s2 := a2[1:3]       //pego uma parte do array, do indice 1 ao 3
	fmt.Println(a2, s2) //slice não é um segundo array, e sim um ponteiro pra parte do array

	s3 := a2[:2] // novo slice apontando do inicio ate o indice 2 (nao inclui o indice 2)
	fmt.Println(s3)

	//Slice de um slice
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
