package main

import (
	"fmt"
	"math"
	"reflect"
)

func main(){
	//números inteiros
	fmt.Println(1,2,1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)...
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal
	i1 := math.MaxInt64	
	fmt.Println ("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// número reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é ", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	//string
	s1 := "Olá meu nome é Karina"
	fmt.Println(s1 + "!")
	fmt.Println ("O tamanho da String é", len(s1))

	//String com multiplas linhas
	s2 := `Ola
	meu
	nome
	é
	Karina`
	fmt.Println("o tamanho da string é ", len (s2))
}