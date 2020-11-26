package main

import (
	"fmt"
	"time"
)

func talk(people, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", people, text, (i + 1))
	}
}

func main() {
	// talk("Maria", "Por que cê num fala comigo?", 3)
	// talk("João", "Só posso falar depois de você?", 1)

	// go talk("Maria", "Ei...", 500)
	// go talk("John", "Fala...", 500)

	go talk("Maria", "Entendi...", 10)
	talk("João", "Parabéns...", 5)
}
