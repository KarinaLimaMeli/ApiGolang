package main

import "fmt"

// ponteiro é uma referencia de memoria
func main (){
	i := 1

	var p *int = nil

	p = &i //pegando o enderenço
	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p,*p, i)
}