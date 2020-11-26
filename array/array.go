package main

import "fmt"

func main() {
	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.9, 4.2, 9.1

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("A média é %.2f\n", media)

	fmt.Println("###################################################")

	numeros := [...]int{1, 2, 3, 4, 5}

	for i, n := range numeros {
		fmt.Printf("%d) %d\n", i, n)
	}

	for _, n := range numeros {
		fmt.Printf("%d\n", n)
	}
}
