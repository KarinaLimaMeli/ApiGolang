package main

import "fmt"

func main() {

	//Com comado Print imprime na mesma linha
	fmt.Print("Mesmo")
	fmt.Print("linha.")
	// Com esse comado Println consigo ter a quebra de linha
	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	// 2 opçoes de impressão 
	fmt.Printf("\n%d %f %.1f %t %s ", a,b,b,c,d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
