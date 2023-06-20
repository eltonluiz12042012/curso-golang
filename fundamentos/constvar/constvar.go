package main

import (
	"fmt"
	m "math"

)

/*
m "math" -> Atribui uma alias ao nome do pacote
_ m "math" -> Define que o pacote foi importado mas não usado
*/
func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * m.Pow(raio, 2)
	fmt.Println("A área de circunferência é", area)

	const (
		A = 1
		B = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(A, B, c, d)

	var f, g bool = false, true

	h, i, j, k := 3.989, true, "epa", 100

	fmt.Println(f, g, h, i, j, k)

}
