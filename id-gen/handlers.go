package main

import (
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func requestID(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//key := r.Form.Get("key")
	//result := Incr(key)

	uuid.New().String()
	//fmt.Println(uuid)
	result := true

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatBool(result)))
}
