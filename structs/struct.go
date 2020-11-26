package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}

	return total
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    2.50,
		desconto: 0.1,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"notebook", 1800.00, 0.1}

	fmt.Println(produto2, produto2.precoComDesconto())

	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtd: 2, preco: 12.10},
			item{produtoID: 2, qtd: 1, preco: 23.90},
			item{produtoID: 11, qtd: 100, preco: 3.90},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
}
