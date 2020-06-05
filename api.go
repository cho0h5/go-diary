package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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