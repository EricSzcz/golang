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

	//sem sinal (só positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal int8 int16 int32 int64
	i1 := math.MaxFloat64
	fmt.Println("O valor máximo de inr é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da variavel i2 é", i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(bo)

	s1 := "Olá meu nome é eric"
	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho de string é", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Eric`
	fmt.Println(s2 + "!!")
	fmt.Println("O tamanho de string é", len(s2))

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
