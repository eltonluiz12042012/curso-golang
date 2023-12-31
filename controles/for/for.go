package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i, ")")
		i++
	}

	for i := 0; i <= 10; i += 2 {
		fmt.Printf("%d - ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par %d:\n", i)
		} else {
			fmt.Printf("Impar %d:\n", i)
		}
	}

	//Loop infinito
	for {
		fmt.Println("Para Sempre", time.Second)
		time.Sleep(time.Second)

	}
}
