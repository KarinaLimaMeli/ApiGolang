package main

import "fmt"

// segundo int serve para referenciar o tipo de retorno, se tornado obrigatorio
// o return
func somar(a int, b int) int{
	return a + b 
}

// função chamada imprimir e não retorna nada
func imprimir(valor int){
	fmt.Println(valor)
}

func main() {
	resultado := somar (3,4)
	imprimir(resultado)
}