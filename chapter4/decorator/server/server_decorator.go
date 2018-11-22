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
	fmt.Fprintf(s.LogWriter, "\nRequest URI.  : %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host          : %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method        : %s\n", r.Method)
	fmt.Fprintf(s.LogWriter, "---------------------------------------")

	s.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintf(w, "Error trying to retrive data from Basic auth\n")
	}
}

func main() {
	//http.Handle("/", &MyServer{})

	/*
		http.Handle("/", &LoggerServer{
			LogWriter: os.Stdout,
			Handler:   &MyServer{},
		})*/

	mySuperServer := &LoggerServer{
		Handler: &BasicAuthMiddleware{
			Handler:  new(MyServer),
			User:     "admin",
			Password: "1234",
		},
		LogWriter: os.Stdout,
	}

	http.Handle("/", mySuperServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
