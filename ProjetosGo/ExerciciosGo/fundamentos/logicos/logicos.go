package main

import "fmt"

// se vc teve 2 empregos vc compra a tv e toma sorvete
// se vc so teve 1 emprego vai comprar uma tv de 32 e toma sorvete
// se vc não trabalhou, não vai fazer nada

func compras(trab1, trab2 bool) (bool,bool,bool)  {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2 
	 return comprarTv50, comprarTv32, comprarSorvete
}

func main () {
	tv50, tv32, sorvete := compras(false,false) // tv50, tv32, sorvete vai receber o resultado de compras
	fmt.Printf("Tv50: %t,Tv32: %t, Sorvete: %t, Saudável: %t",
	tv50, tv32, sorvete, !sorvete)
}