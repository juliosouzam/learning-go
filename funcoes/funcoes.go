package main

import "fmt"

func trocar(a, b int) (p, s int) {
	p = b
	s = a
	return
}

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

var sum = func(a, b int) int {
	return a + b
}

var multi = func(a, b int) int {
	return a * b
}

var exec = func(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func media(n ...float64) float64 {
	total := 0.0

	for _, valor := range n {
		total += valor
	}

	return total / float64(len(n))
}

func imprimirAprovadors(aprovados ...string) {
	fmt.Println("LISTA DE APROVADOS")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i, aprovado)
	}
}

func closure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialOld := fatorial(n - 1)
		return n * fatorialOld
	}
}
