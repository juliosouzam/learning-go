package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<p>Boa noite.
		<a href="/horaCerta">clique aqui</a> pra vê a hora atual</p>
	`)
}

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/horaCerta", horaCerta)
	fmt.Println("Exec...")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
