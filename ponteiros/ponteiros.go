package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i

	i++
	*p++

	fmt.Println(i, *p)

	if true && false || (false && true) {
		fmt.Print("")
	}

	nota := 6
	switch nota {
	case 6:
		fmt.Println("passô")
		fallthrough
	case 5:
		fmt.Println("A")
	}

	n := 1

	inc1(n)
	fmt.Println(n)

	inc2(&n)

	fmt.Println(n)

}

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}
