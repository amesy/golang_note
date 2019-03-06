package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>result: %s</h2>", r.URL)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("user"))
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/", userHandler)
	http.ListenAndServe(":8080", nil)
}
