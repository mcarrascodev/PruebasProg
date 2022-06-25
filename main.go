package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Println("Server Runing")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo en MercadoLibre")

}
