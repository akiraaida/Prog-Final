package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
        "io"
)

func parse(data io.Reader) {

    tokenStream := html.NewTokenizer(data)

    for {
        token := tokenStream.Next()
        if token == html.ErrorToken {
            return
        }
                
        if token == html.StartTagToken {
            checkToken := tokenStream.Token()
            
            if(checkToken.Data == "p"){
                checkText := tokenStream.Next()
                if checkText == html.TextToken {
                    byteData := tokenStream.Text()
                    fmt.Println(string(byteData[:]))
                }
            }
        }
    } 
}

func retrieve(site string) {
        
    // Get takes a string and returns a response and error
    // The error is nil if the response was successful
    // The response is a data structure that contains alot
    // of information about the page. The body being the
    // main part that is needed
    resp, err := http.Get(site)
    if err != nil {
    	return
    }
    // The response body must be closed due to potential leaking of
    // open files
    defer resp.Body.Close()
        
    // Calls parse with the body that was extracted from the Get
    // The body is of type ReadCloser which uses the interface
    // io.Reader
    parse(resp.Body)
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

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
    	log.Fatal(err)
    }
}
