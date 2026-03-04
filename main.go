package main

import (
	"fmt"
	"net/http"
	"html/template"
)

// Plans are made to be broken
// - figure out some interactivity (forms)
// - get a database working
// - figure out what we actually want to do

type Palabra struct {
    Word string
}
type PageContent struct {
	Title string
    Words []Palabra
}

func getMain(w http.ResponseWriter, r * http.Request) {
	fmt.Println("Accessing main page")

    words := []Palabra{Palabra{"Hi"}, Palabra{"Hello"}}
    thing := PageContent{Title: "Home", Words: words}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, thing)
}

func main() {
	http.HandleFunc("/", getMain)

	fmt.Println("Listening at 8080")

	styles := "/static/css"

    thing := http.FileServer(http.Dir(styles[1:]))
	strip := http.StripPrefix(styles, thing)
    http.Handle("/static/css/", strip)
	http.ListenAndServe(":8080", nil)
}
