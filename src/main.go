package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./index")))

	err := http.ListenAndServe("www.akiraaida.me:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
