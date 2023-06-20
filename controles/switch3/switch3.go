package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "booleano"
	default:
		return "não sei"

	}
}

type Pessoa struct{}

func main() {
	var p Pessoa
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo("Teste"))
	fmt.Println(tipo(1))
	fmt.Println(tipo(1.1))
	fmt.Println(tipo(true))
	fmt.Print(tipo(p))

}
