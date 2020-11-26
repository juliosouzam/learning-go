package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json

	p1 := produto{ID: 1, Nome: "Notebook", Preco: 1800.9, Tags: []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct

	var p2 produto
	jsonString := `{"id":2,"nome":"Mouse","preco":189.9,"tags":["Promoção","Eletrônico"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
