package main

import (
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	time.Sleep(3 * time.Millisecond)
}
