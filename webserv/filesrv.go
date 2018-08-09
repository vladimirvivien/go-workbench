package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	fmt.Println("serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./"))))
}
