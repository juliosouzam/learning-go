package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Depois lido")
}

func main() {
	c := make(chan int) // channel without buffer

	go rotina(c)

	fmt.Println(<-c)        // operação bloqueante
	fmt.Println("Foi lido") // operação bloqueante
	fmt.Println(<-c)        // operação bloqueante
	fmt.Println("Fim")      // operação bloqueante
}
