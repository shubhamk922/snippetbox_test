package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
1 Handler - func (w http.ResponseWriter, r *http.Request)
2 Router -- mux.NewServerMux and then map the route with the handler
3 Web Server create a server - ListenAndServer , In go you dont need a third party server like apache or ngnix
*/

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

func main() {
	// create a new server mux and map the / to home handler
	// http.Handler is an interface have function handleFunc
	// http.Server mux is a interface type of http Handler
	/*for the sake of clarity, maintainability and security, it’s generally a good idea to avoid http.DefaultServeMux and the corresponding helper functions. Use your own locally-scoped servemux instead, like we have been doing in this project so far.*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("GET /snippet/view/{id}", snippetview)
	mux.HandleFunc("POST /snippet/view", snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	//err = http.ListenAndServe(":4000", nil) This will use http.DefaultServerMux
	if err != nil {
		log.Fatal("Issue is starting sever at port 4000")
	}
}
