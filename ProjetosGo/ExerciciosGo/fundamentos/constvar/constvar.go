package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 31415
	var raio = 3.2 // tipo (folat64) imferido pelo compliador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circuferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// tipagem
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
