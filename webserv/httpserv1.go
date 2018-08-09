package main

import "fmt"
import "net/http"

// creates server and starts
func main() {
	handler := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Hello from Go!")
	}

	fmt.Println("running on port 4040")
	http.ListenAndServe(":4040", http.HandlerFunc(handler))
}
