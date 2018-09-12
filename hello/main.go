package main

import (
	"log"
	"net/http"
)

// $ go install workshop/hello
func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(hello)); err != nil {
		log.Println(err)
	}
}
