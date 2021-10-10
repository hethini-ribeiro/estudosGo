package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //gera um numero aleatorio, dando o nano segundo da data atual
	r := rand.New(s)
	return r.Intn(10) //gera um número aleatorio ate 10
}

func main() {
	//estou inicializando uma variavel no if (como no laço For)
	if i := numeroAleatorio(); i > 5 { // tbm suportado no switch
		fmt.Println("Ganhou:", i)
	} else { //o else precisa estar na mesma linha
		fmt.Println("Perdeu:", i)
	}
}
