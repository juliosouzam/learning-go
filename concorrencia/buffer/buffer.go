package main

import "fmt"

func rotine(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("Exec!")
	c <- 5
	c <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotine(ch)

	fmt.Println(<-ch)
}
