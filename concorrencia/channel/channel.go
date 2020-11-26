package main

import (
	"fmt"
	"time"
)

func times(b int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * b

	time.Sleep(time.Second)
	c <- 3 * b

	time.Sleep(3 * time.Second)
	c <- 4 * b
}

func main() {
	// ch := make(chan int, 1)

	// ch <- 1 // enviando dados para o canal
	// <-ch    // recebendo dados
	// ch <- 2
	// fmt.Println(<-ch)
	c := make(chan int)

	go times(2, c)

	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println(<-c)
}
