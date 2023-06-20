package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"

)

func alteraMensagem(s *string, nova string) {
	*s = nova
}

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(320000))

	//Tipos sem sinal -> uint8 uint16 uint32 uint64

	var b byte = 255 //Alias para um tipo uint8

	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = '\\' //representa o caractere na tabela unicade (int32)
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))

	// Float
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//Boolean

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	mensagem := "Olá, meu nome é Elton Luiz"
	alteraMensagem(&mensagem, "Elton Luiz dos Santos do Franco")
	fmt.Println(mensagem)
	fmt.Println(len(mensagem))
	fmt.Println(strings.Replace(mensagem, " ", "-", len(mensagem)))
	fmt.Println(strings.SplitAfter(mensagem, " ")[0:2])
	fmt.Println(strings.ToTitle(mensagem))
	//mensagem2 := strings.Clone(mensagem)

}
