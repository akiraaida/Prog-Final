<h3 style="text-align:center">Programming Languages - Final Project</h3>
<h4 style="text-align:center">Akira Aida - 100526064</h4>
***
<p><b>1. Github:</b> Your code just be in a public repository on github.com.<br>
<p><b>- Repository:</b> My public repository can be found at https://github.com/akiraaida/Prog-Final.</p>
***
<p><b>2. Activities:</b> You must make frequent updates to the git repo.  This will be verified using “git log” command.  You must have an entry to your git repo at least every two days.

<b>- Commit Log:</b> I have made frequent updates to my github repository within the past few weeks. However they have not all been within a two period due to other commitments.
***
<p><b>3. Problem Statement and Language Selection:</b> You must clearly state the problem you are addressing, and the language of your selection to solve the problem.

<b>- Problem Statement:</b> The problem at hand is finding common words within an HTML page without the reliance of a user's hardware. This is done through doing the calculations on a server (instead of client side) and by web scraping the HTML. This would allow the user to find relevant words on the topic to branch off of. Or allow for a very basic search engine indexing algorithm (of course this method would be exploitable through means of hiding text in HTML documents).

<b>- Language Selection:</b> The language I chose to use for this problem was the Go programming language. The Go programming language has built in web server support which solves the problem of not relying on the user's hardware. As well as a package already written for navigating the DOM (this package is golang.org/x/net/html) which can easily be installed with "go get golang.org/x/net/html".
***
<p><b>4. A Brief Survey of Alternatives:</b> State how the same problem can be solved by some other languages.  You don’t have to implement a solution using a second language, but you are expected to know about them.

<b>- Python:</b> Python has the SimpleHTTPServer module which allows for basic server hosting. Python also has many libraries that make webscraping alot easier. Some notable ones are BeautifulSoup4 which allows easy navigation of the DOM. Another one being Mechanize which allows for selecting attributes in HTML forms and submitting them similar to how a browser does so.

<b>- NodeJS:</b> Allows for hosting a server using Javascript. Allows for code re-use between front and backend (to a degree) and allows the comfort of only using one language. Has many packages that make web parsing easier such as Request which makes HTTP calls cleaner and Cheerio to help traverse the DOM.
***
<p><b>5. Build Tools:</b> Describe the build tools required to develop the solution in your language.  You should have a Makefile whenever possible in your git repo.

<b>- Build Tools:</b> The build tools required to develop the solution for my problem are solely the Makefile. As well as having Go 1.6 installed with the GOPATH set.
***
<p><b>6. Language Features:</b> Describe by means of code walk the language features used in your solution.  You are expected to highlight the most impactful usage of some language features.

<b>Highlights:</b>

http.Handle("/", http.FileServer(http.Dir("./home")))

http.HandleFunc("/submit", handleSubmit)

err := http.ListenAndServe(":8080", nil)

resp, err := http.Get(site)

tokenStream := html.NewTokenizer(data)

token := tokenStream.Next()
***
<p><b>7. Relating to the Course:</b> Continuing with (6), you should refer back to the topics covered in the course including type systems, type inference, lexical scoping and closure, functions as data, coroutines, list comprehension, etc.  You can also relate back to Scala, Clojure, Java whenever appropriate to illustrate the similarity and distinction of the language of your choice.<br>

<b>- Type System:</b> Minimalistic approach with primitive types such as booleans, numerics (float32-float64, (u)int8-(u)int64, and more), and strings. As well as composite types like arrays, slices, maps, and more. In Go, structs are the way to create concrete user-defined types with interfaces being requirements for a type.

<b>- Type Inference:</b> Go's type inference works by inferring the type from the value given. (i := 5 is an int). Types can also be specified by doing a syntax of "var i int".

<b>- Lexical Scoping and Closure:</b> The scope of the Go programming language is that where the variable is defined in the program is where is can be used functionally. Meaning if it's defined in a for loop then it can't be used outside of it. The variable needs to be defined in a function body and if a variable is defined again inside of an inner statement, then the outer variable definition is shadowed. An example of this is...

func main () {

&emsp;val := 10

&emsp;if true {

&emsp;&emsp;val := 5

&emsp;&emsp;fmt.Println(val) //Prints 5

&emsp;}

&emsp;fmt.Println(val) // Prints 10

}
***
<p><b>8. Live Demo</b></br>
