package main

import (
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/post/{code}", logging(getPost)).Methods("GET")
	router.HandleFunc("/post", logging(postPost)).Methods("POST")
	router.HandleFunc("/posts", logging(getPosts)).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":3000", router)

	// err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", router)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
