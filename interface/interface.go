package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	p := pessoa{nome: "Júlio", sobrenome: "César"}
	fmt.Println(p.toString())
	imprimir(p)

	coisa := produto{nome: "Calça", preco: 129.90}

	imprimir(coisa)

	f := ferrari{modelo: "F40", turboLigado: false, velocidadeAtual: 0}
	f.ligarTurbo()

	var f2 esportivo = &ferrari{modelo: "F40", turboLigado: false, velocidadeAtual: 0}
	f2.ligarTurbo()

	fmt.Println(f, f2)
}
