package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Printf("Aprovado nota %.3f", nota)
	} else {
		fmt.Printf("Reprovado nota %.2f", nota)
	}
}

func main() {
	imprimirResultado(7.999)
}
