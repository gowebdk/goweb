package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	files := http.FileServer(http.Dir("/home"))
	mux.Handle("/home", http.StripPrefix("/home/", files))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, welcome , was waiting for you to make me alive...")
}
