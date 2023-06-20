package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Data 1:", d1)
	fmt.Println("Data 2:", d2)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		nome string
	}

	p1 := Pessoa{nome: "João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoa:", p1 == p2)      //Compara valores
	fmt.Println("Endereços:", &p1 == &p2) //Compara endereços de memória dos objetos
}
