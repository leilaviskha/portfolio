package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Serveur lancé sur http://localhost:8055")
	http.ListenAndServe(":8055", nil)
}