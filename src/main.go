package main

import (
	"fmt"
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("website:", r.Form["website"])
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./index")))
	http.HandleFunc("/search", search)

	err := http.ListenAndServe("www.akiraaida.me:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
