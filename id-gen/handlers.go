package main

import (
	"net/http"
	"strconv"
)

func requestID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	result := Incr(key)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatBool(result)))
}
