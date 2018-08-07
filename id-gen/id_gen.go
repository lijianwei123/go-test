package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/id_gen", requestID)

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
