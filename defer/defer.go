package main

import "fmt"

func get(n int) int {
	defer fmt.Println("FIM!")
	if n > 5000 {
		fmt.Println("Maior que 5000")
		return 5000
	}

	fmt.Println("Menor que 5000")
	return n
}

func main() {
	fmt.Println(get(6000))
	fmt.Println(get(3000))
}
