package main

import "fmt"

type nota float64

func (n nota) entre(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func notaPraConceiro(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	}
	if n.entre(7.0, 8.99) {
		return "B"
	}
	if n.entre(5.0, 6.99) {
		return "C"
	}
	if n.entre(3.0, 4.99) {
		return "D"
	}

	return "E"
}

func main() {
	fmt.Println(notaPraConceiro(7.9))
	fmt.Println(notaPraConceiro(9.9))
	fmt.Println(notaPraConceiro(5.9))
	fmt.Println(notaPraConceiro(1.9))
}
