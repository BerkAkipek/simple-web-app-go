package main

import (
	"fmt"
	"net/http"

	"github.com/BerkAkipek/simple-web-app-go/pkg/handlers"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server listening on http://localhost%v/\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
