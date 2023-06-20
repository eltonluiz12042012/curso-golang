package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	//fmt.Println("O valor de x é: " + x) -> isso não funciona em golang

	xs := fmt.Sprint(x) //Retorna o valor do parâmetro passado como tipo string

	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é: ", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.2f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
	/*
	 %d - Número Inteiros
	 %f - Número de ponto flutuante
	 %.2f - Número de ponto flutuante com duas casas decimais
	 %t - Valores booleanos
	 %s - Valores do tipo string
	 %v - Curinga imprime todos os tipos
	*/
}
