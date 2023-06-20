package main

import (
	"fmt"
	"strconv"

)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste" + string(92)) //Imprime o caractere 92 da tabela ascii

	fmt.Println("Convertendo de int para string" + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //1- true / 2 - false
	if b {
		fmt.Printf("Verdadeiro: %t", b)
	}
}
