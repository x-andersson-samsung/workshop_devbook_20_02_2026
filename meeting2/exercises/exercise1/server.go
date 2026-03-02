package main

import (
	"fmt"
	"io"
	"net/http"
)

type Server struct{}

func (s Server) handleGet(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("GET request")
	w.Write([]byte("Hello World!"))
}

func (s Server) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("POST request with body: " + string(body))
	w.Write(body)
}

func main() {
	var s Server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo", s.handleGet)
	mux.HandleFunc("POST /echo", s.handlePost)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
