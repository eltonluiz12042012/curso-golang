package main

import "fmt"

func pegaEndereco(valor *int) *int {
	return valor
}

func main() {
	/*a := 1
	fmt.Println(&a)
	var b *int
	b = &a
	fmt.Println("Valor a antes da alteração", *b)
	fmt.Println("O endereco de b", &b)
	*b = 10
	fmt.Println(a)
	fmt.Println(*b)*/
	i := 1
	var p *int = nil
	p = &i
	*p++
	i++
	//fmt.Println(p, *p, i, &i)

	idade := 12
	fmt.Println(&idade)

	var novo *int
	novo = pegaEndereco(&idade)
	fmt.Println(novo)

	*novo = 43
	fmt.Println(idade)
	fmt.Println(*novo)

}
