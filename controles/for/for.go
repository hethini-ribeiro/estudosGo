package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 { //não tem while em go
		fmt.Println(i)
		i++
	}
	fmt.Println("*******************")
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}
}
