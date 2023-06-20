package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3

	fmt.Println(i)

	x, y := 1, 2
	fmt.Print(x, y)

	x, y = y, x
	fmt.Print(x, y)
}
