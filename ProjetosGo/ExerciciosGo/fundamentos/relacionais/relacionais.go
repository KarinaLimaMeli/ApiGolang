package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("String:", "Banana" == "Banana") // verificando se são iguais 
	fmt.Println("!=", 3 != 2) // verificando se é diferente 
	fmt.Println("<", 3 < 2) // verificando se é menor ou maior
	fmt.Println(">", 3 <= 2) // verificando se é menor ou maior
	fmt.Println("<=", 3 <= 2) // verificando se é menor ou igual
	fmt.Println(">=", 3 >= 2) // verificando se é menor ou maior

	//objeto tipo data 
	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string 
	}
		p1 := Pessoa{"João"}
		p2 := Pessoa{"João"}
		fmt.Println("Pessoas:", p1 == p2)
		
	}