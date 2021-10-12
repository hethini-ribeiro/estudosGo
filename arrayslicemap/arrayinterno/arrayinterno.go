package main

import "fmt"

func main() {
	//Provando que os slice fazem referencia ao mesmo trecho de memoria

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)
	// Atribuindo valor novo apenas ao slice 1 altera o mesmo indice e valor do slice 2
	s1[0] = 7
	fmt.Println(s1, s2)
}
