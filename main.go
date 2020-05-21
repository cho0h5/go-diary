package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Data struct {
	Code        string
	Title       string
	Description string
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go writeLog(r.URL.Path)
		f(w, r)
	}
}

func writeLog(a ...interface{}) {
	file, err := os.OpenFile("request.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(a)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	i := Data{Code: p["code"], Title: "title", Description: "desc"}
	json.NewEncoder(w).Encode(i)
}

func postPost(w http.ResponseWriter, r *http.Request) {
	fileData := r.FormValue("article")
	go func(fileData string) {
		path := "data/"
		os.MkdirAll(path, os.ModePerm)

		fileName := time.Now().Format("20060102_150405")
		file, err := os.Create(path + fileName + ".txt")
		if err != nil {
			writeLog(err)
		}
		defer file.Close()

		fmt.Println(fileData) //
		file.WriteString(fileData)
	}(fileData)
	// fmt.Fprint(w, "postPost test")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./data/")
	if err != nil {
		writeLog(err)
	}
	for _, f := range files {
		fmt.Fprintln(w, f.Name())
	}
}

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
