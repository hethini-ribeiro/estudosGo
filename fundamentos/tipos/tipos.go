package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivo)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Byte é um", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O valor máximo do int é", reflect.TypeOf(i1))

	var i2 rune = 'a' //represeta um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numero reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome é Hethini"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//String com multiplas linhas
	s2 := `Ola 
	meu 
	nome 
	é 
	Hethini`
	fmt.Println("O tamanho da string é", len(s2))
}
