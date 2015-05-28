package main

import (
	"fmt"
	"net/http"
)

type hello struct {}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!")
}

func main() {
	var h hello
	http.ListenAndServe("localhost:4000", h)
}
