package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola desde go")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Seervidor corriendo en http//localhost:8080")
	http.ListenAndServe(":8080", nil)
}
