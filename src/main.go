package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleSearch(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view.html")
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./home")))
	http.HandleFunc("/search/", handleSearch)

	err := http.ListenAndServe("www.akiraaida.me:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
