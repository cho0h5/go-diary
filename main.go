package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Data struct {
	Code        string
	Title       string
	Description string
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go func() {
			file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			log.SetOutput(file)
			log.Println(r.URL.Path)
		}()
		f(w, r)
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	i := Data{Code: p["code"], Title: "title", Description: "desc"}
	json.NewEncoder(w).Encode(i)
}

func postPost(w http.ResponseWriter, r *http.Request) {
	// p := mux.Vars(r)
	// i := Data{Code: p["code"], Title: "title", Description: "desc"}
	// json.NewEncoder(w).Encode(i)
	fmt.Fprint(w, "postPost test")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/post/{code}", logging(getPost)).Methods("GET")
	router.HandleFunc("/post}", logging(postPost)).Methods("POST")
	router.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":3000", router)
	// err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", router)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
