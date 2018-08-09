package main

import "fmt"
import "net/http"

// creates server and starts
func main() {
	fmt.Println("running on port 4040")
	http.ListenAndServe(
		":4040",
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			fmt.Fprint(resp, "Hello from Go!")
		}),
	)
}
