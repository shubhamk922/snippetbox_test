package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Home"))
}

func snippetview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if id < 1 || err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to Snippet View %d...", id)
	//w.Write([]byte(fmt.Sprintf("Welcome to Snippet View %d", id)))
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Snippet Create"))
}

func snippetCreatePost(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Server", "Go") // Add all haeder maps

	w.WriteHeader(http.StatusCreated) // status code should be set before writing response

	// if you see writing response body there is a method Write that means
	// http.ResponseWriter has Write method ie it implement io.Writer interface
	// fmt.Fprint( w, "Hello")
	// io.WriteString(w, "Hello")
	w.Write([]byte("Save a new snippet")) // write the response
}
