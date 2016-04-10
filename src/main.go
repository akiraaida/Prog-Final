package main

import (
	//	"fmt"
	//	"golang.org/x/net/html"
	//	"io/ioutil"
	"log"
	"net/http"
)

/*
func whitespace(item string)(res bool) {
    ws := []string{"", " ", "\t", "\n"}
    for i := 0; i < len(ws); i++ {
        if item == ws[i] {
            res = true
            return
        }
    }
    res = false
    return
}

func parse(data string) {

    for i := 0; i < len(data); i++ {
        ws := whitespace(string(data[i]))
        if ws == false {
            fmt.Println(string(data[i]))
        }
    }
}
*/

func retrieve(site string) {
	resp, err := http.Get(site)
	if err != nil {
		return
	}
	body := resp.Body
	defer body.Close()
	//parse(string(body))
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		site := r.FormValue("website")
		retrieve(site)
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./home")))
	http.HandleFunc("/submit", handleSubmit)

	err := http.ListenAndServe("www.akiraaida.me:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
