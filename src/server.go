package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
)

func partition(counts []int, start int, end int) (mid int) {

	// Move larger items before the pivot and smaller after it
	pivot := counts[end]
	i := start - 1

	for j := start; j < end; j++ {
		if counts[j] >= pivot {
			i = i + 1
			temp := counts[i]
			counts[i] = counts[j]
			counts[j] = temp
		}
	}

	temp := counts[i+1]
	counts[i+1] = counts[end]
	counts[end] = temp

	mid = i + 1

	return
}

func quickSort(counts []int, start int, end int, channel chan []int) { //(sortedCounts []int) {

	// Continuously partition until start < end
	if start < end {
		mid := partition(counts, start, end)
		quickSort(counts, start, mid-1, channel)
		quickSort(counts, mid+1, end, channel)
	}
	sortedCounts := counts
	//return
	channel <- sortedCounts
}

func sortMap(w http.ResponseWriter, wordMap map[string]int) {

	// Create a map to hold the count, word equivalent
	tempMap := make(map[int][]string)
	var counts []int

	// Go through the word map and create the count, word
	// equivalent
	for word, count := range wordMap {
		tempMap[count] = append(tempMap[count], word)
	}

	// Go through the count, word map and append the counts
	for count := range tempMap {
		counts = append(counts, count)
	}

	// Quicksort the counts slice to be sorted from most occuring to least occuring
	// Just for the sake of learning, created a thread and channel for the quicksort
	channel := make(chan []int)
	go quickSort(counts, 0, len(counts)-1, channel)
	counts = <-channel
	//counts = quickSort(counts, 0, len(counts)-1)

	fmt.Fprintf(w, "<table class=\"table table-bordered\"><thead><tr><th>Rank</th><th>Word</th><th>Occurence</th></tr></thead><tbody>")

	// Iterate through counts, throwing away the index
	num := 1
	for _, count := range counts {
		// Get the word where the count specifies it should be
		for _, word := range tempMap[count] {
			// Print the data in descending order
			fmt.Fprintf(w, "<tr>")
			fmt.Fprintf(w, "<td><b>%d</b></td><td><b>%s</b></td><td><b>%d</b></td>", num, word, count)
			fmt.Fprintf(w, "</tr>")

			num++
			if num == 16 {
				break
			}
		}
		if num == 16 {
			break
		}
	}
	fmt.Fprintf(w, "</tbody></table>")
}

func countOcc(w http.ResponseWriter, wordList []string) {

	// Create a word map to count the occurences of the words
	wordMap := make(map[string]int)

	// Loop through every word and normalize them to lower case
	// As well as removing characters that throw off the count
	for i := 0; i < len(wordList); i++ {

		wordList[i] = strings.Replace(wordList[i], ".", "", -1)
		wordList[i] = strings.Replace(wordList[i], ",", "", -1)
		wordList[i] = strings.Replace(wordList[i], "(", "", -1)
		wordList[i] = strings.Replace(wordList[i], ")", "", -1)
		wordList[i] = strings.Replace(wordList[i], "\"", "", -1)
		wordList[i] = strings.ToLower(wordList[i])

		if wordList[i] != "" {
			wordMap[wordList[i]]++
		}
	}
	sortMap(w, wordMap)
}

func parse(w http.ResponseWriter, data io.Reader) {

	// HTML is tokenized using a tokenizer
	// on the response body that was fetched
	tokenStream := html.NewTokenizer(data)

	// Create a slice for the words that will be
	// parsed from the HTML
	wordList := make([]string, 0)

	// Loops through all of the tokens one by one
	for {
		// Gets the next token
		token := tokenStream.Next()

		// If the next token is an error token
		// like the end of file then it breaks
		// out of the loop
		if token == html.ErrorToken {
			break
		}

		// If the token is a start tag
		if token == html.StartTagToken {

			// Then check if the tag is a <p> tag
			checkToken := tokenStream.Token()
			if checkToken.Data == "p" {
				depth := 1
				for depth > 0 {

					checkNext := tokenStream.Next()

					switch {
					// Check if it's a text token, if so then parse it
					case checkNext == html.TextToken:
						// Get the text as a byte array
						byteData := tokenStream.Text()
						// Initialize a variable that will hold the words
						// parsed from byte -> string conversion
						word := ""

						// Loop through the byte array, concatenating to the
						// word string unless it is a 32 (space), 10 (newline) or 9 (tab).
						// If it is one of those and the word is not empty, then a word
						// has been found and append it to the slice created at the start
						for i := 0; i < len(byteData); i++ {
							if byteData[i] != 32 && byteData[i] != 10 && byteData[i] != 9 {
								word += string(byteData[i])
							} else {
								if word != "" {
									wordList = append(wordList, word)
									word = ""
								}
							}
						}

						// If we make it to the end of the byte array and there was no
						// 32, 10, or 9 to signify the end of a the word then we add
						// it to the slice
						if word != "" {
							wordList = append(wordList, word)
						}

					// Check if it's a start tag within the <p> tag, if so increase
					// the depth so the data can be either skipped (if non text) or
					// parsed
					case checkNext == html.StartTagToken:
						depth++
					// When the depth is 0 then the matching <p> tag has been found
					case checkNext == html.EndTagToken:
						depth--
					}
				}
			}
		}
	}
	countOcc(w, wordList)
}

func retrieve(w http.ResponseWriter, site string) {

	// Get takes a string and returns a response and error
	// The error is nil if the response was successful
	// The response is a data structure that contains alot
	// of information about the page. The body being the
	// main part that is needed
	resp, err := http.Get(site)
	// If no input or an error, don't proceed further
	if err != nil {
		return
	}
	// The response body must be closed due to potential leaking of
	// open files
	defer resp.Body.Close()

	// Calls parse with the body that was extracted from the Get
	// The body is of type ReadCloser which uses the interface
	// io.Reader
	parse(w, resp.Body)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {

	// The handler for the HTML form submission.
	// Takes the inputted data and passes it along
	fmt.Fprintf(w, "<html><head><link rel=\"stylesheet\" type=\"text/css\" href=\"index.css\"><link rel=\"stylesheet\" type=\"text/css\" href=\"bootstrap.min.css\"></head><body>")
	fmt.Fprintf(w, "<a id=\"homeButton\" href=\"/\">Home</a><p>")
	if r.Method == "POST" {
		r.ParseForm()
		site := r.FormValue("website")
		retrieve(w, site)
	}
	fmt.Fprintf(w, "</body></html>")
}

func main() {

	// Handle responds to incoming HTTP requests
	// FileServer returns a handler that serves any HTTP requests
	// with the contents of the directory specified
	http.Handle("/", http.FileServer(http.Dir("./home")))

	// HandleFunc connects the function and the path URL
	http.HandleFunc("/submit", handleSubmit)

	// ListenAndServe listens on the network address
	// specified with a nil handler (specifying to use
	// the default)
	// Creates a new go routine for every connection
	err := http.ListenAndServe("www.akiraaida.me:80", nil)

	// If an error occurs while serving, outputs the
	// error and then exits then closes the server
	if err != nil {
		log.Fatal(err)
	}
}
