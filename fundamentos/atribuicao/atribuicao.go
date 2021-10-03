package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferencia de tipo
	i += 3 // i = i+3
	i -= 3 // i = i-3
	i /= 3 // i = i/3
	i *= 3 // i = i*3
	i %= 3 // i = i%3 (MOD - resto)

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x // inverte/troca os valores da variaveis
	fmt.Println(x, y)

}
