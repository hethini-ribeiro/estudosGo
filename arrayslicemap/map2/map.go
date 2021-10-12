package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.66,
		"Pedro Júnior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00
	delete(funcsESalarios, "Inecistente") // tentando excluir um indice q n existe e não da erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
