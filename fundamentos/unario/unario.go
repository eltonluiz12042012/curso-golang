package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	// ++x Isso não existe em golang

	y++

	//fmt.Println(x == y--) isso tb nã funciona
	y--
	x = y
	fmt.Println(x)
}
