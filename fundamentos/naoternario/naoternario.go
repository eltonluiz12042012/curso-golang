package main

import "fmt"

//Em go não existe operação unária
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	//nota>=6?"Aprovado":"Reprovado" - Isso não existe em go
}

func main() {
	fmt.Println(obterResultado(6.4))

}
