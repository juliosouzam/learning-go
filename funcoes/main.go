package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Initializing...")
}

func main() {
	result := somar(1, 2)
	imprimir(result)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1 == d2)

	type Pessoa struct {
		nome string
	}

	x, y := trocar(2, 3)
	fmt.Println(x, y)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sum(2, 3))
	fmt.Println(sub(2, 3))

	resul := exec(multi, 4, 8)

	fmt.Println(resul)

	fmt.Printf("%.2f\n", media(7.7, 8.1, 5.9, 9.9))

	aprovados := []string{"Maria", "João", "Rafael", "Julio"}
	imprimirAprovadors(aprovados...)

	x1 := 20
	fmt.Println(x1)
	printX := closure()
	printX()

	res1 := fatorial(5)

	fmt.Println(res1)
}
