package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	pe := pessoa{nome: "Júlio", sobrenome: "César"}
	fmt.Println(pe, pe.getNomeCompleto())

	pe.setNomeCompleto("Souza César")
	fmt.Println(pe, pe.getNomeCompleto())
}
