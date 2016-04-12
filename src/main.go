package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func parse(data string) {

	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				//if a.Key == "href" {
				fmt.Println(a.Val)
				break
				//}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func retrieve(site string) {
	resp, err := http.Get(site)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	parse(string(body))
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
