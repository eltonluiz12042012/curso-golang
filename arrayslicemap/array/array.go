package main

import "fmt"

func main() {
	var notas [3]float64
	var media, total float64 = 0, 0
	notas[0], notas[1], notas[2] = 7.8, 5.5, 9.0
	for _, v := range notas {
		total += v
	}
	media = total / float64(len(notas))
	fmt.Printf("A Média final é %.2f\n", media)

}
