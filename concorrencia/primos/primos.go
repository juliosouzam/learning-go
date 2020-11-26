package main

import (
	"fmt"
	"time"
)

func isPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	start := 2

	for i := 0; i < n; i++ {
		for primo := start; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				start = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 100)

	go primos(cap(c), c)
	for i := range c {
		fmt.Printf("%d ", i)
	}

	fmt.Println("end")
}
