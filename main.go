package main

import (
	"log"
	"net/http"
)

/*
1 Handler - func (w http.ResponseWriter, r *http.Request)
2 Router -- mux.NewServerMux and then map the route with the handler
3 Web Server create a server - ListenAndServer , In go you dont need a third party server like apache or ngnix
*/

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Home"))
}

func snippetview(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Snippet View"))
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Snippet Create"))
}

func main() {
	// create a new server mux and map the / to home handler
	// http.Handler is an interface have function handleFunc
	// http.Server mux is a interface type of http Handler
	/*for the sake of clarity, maintainability and security, it’s generally a good idea to avoid http.DefaultServeMux and the corresponding helper functions. Use your own locally-scoped servemux instead, like we have been doing in this project so far.*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetview)

	err := http.ListenAndServe(":4000", mux)
	//err = http.ListenAndServe(":4000", nil) This will use http.DefaultServerMux
	if err != nil {
		log.Fatal("Issue is starting sever at port 4000")
	}
}
