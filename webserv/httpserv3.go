package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Make  string
	Model string
}

func main() {
	mux := http.NewServeMux()
	ford := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "application/json")

		car := Car{"Ford", "F150"}
		enc := json.NewEncoder(resp)
		if err := enc.Encode(&car); err != nil {
			fmt.Println(err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	mux.HandleFunc("/ford", ford)
	fmt.Println("started server on port 4040")
	http.ListenAndServe(":4040", mux)
}
