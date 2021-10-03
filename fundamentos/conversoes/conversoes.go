package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste " + string(97)) //printa o caractere da tabela ascii e não o valor numerico

	//int to string
	fmt.Println("Teste " + strconv.Itoa(97))

	//string to int
	num, _ := strconv.Atoi("97")
	fmt.Println(num - 96)

	//conversão de string pra bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}
