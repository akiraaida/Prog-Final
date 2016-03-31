package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./home")))
	http.Handle("/search", http.FileServer(http.Dir("./search")))
	http.Handle("/contact", http.FileServer(http.Dir("./contact")))

	err := http.ListenAndServe("www.akiraaida.me:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
