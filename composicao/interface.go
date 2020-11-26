package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// add other methods
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("turbo ligado")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("baliza")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
