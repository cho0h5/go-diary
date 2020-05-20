package main

import (
	"encoding/json"
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

func getData(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	i := Data{Code: p["code"], Title: "title", Description: "desc"}
	json.NewEncoder(w).Encode(i)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test/{code}", logging(getData)).Methods("GET")

	http.ListenAndServe(":8080", router)
	// err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", router)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
