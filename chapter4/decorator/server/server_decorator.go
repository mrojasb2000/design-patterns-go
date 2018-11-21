package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI.  : %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host          : %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method        : %s\n", r.Method)
	fmt.Fprintf(s.LogWriter, "---------------------------------------")

	s.Handler.ServeHTTP(w, r)
}

func main() {
	//http.Handle("/", &MyServer{})

	http.Handle("/", &LoggerServer{
		LogWriter: os.Stdout,
		Handler:   &MyServer{},
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
