package main

import "fmt"

func main() {
	i := 1

	//Go não tem operações aritméticas de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variavel i
	*p++
	i++
	fmt.Println(p, *p, i, &i)

}
