package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println(s)

	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) //Quando o slice chega na sua capacidade máxima ele dobra o seu tamanho
	fmt.Println(s, len(s), cap(s))
}
