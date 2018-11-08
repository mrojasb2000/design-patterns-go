package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	Msg string
}

// Implement a ServeHTTP method from Handler interface package net/http
func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// server := &MyServer{
	// 		Msg: "Hello, World",
	// 	}
	//
	// 	http.Handle("/", server)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
