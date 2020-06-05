package main

import (
	"net/http"
	"log"
	"os"
)
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