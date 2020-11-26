package main

import (
	"fmt"

	"github.com/juliosouzam/html"
)

func encaminhar(o <-chan string, d chan string) {
	for {
		d <- <-o
	}
}

// multiplexar - misturar mensagens em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Title("https://google.com", "https://youtube.com"),
		html.Title("https://amazon.com", "https://linkedin.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
