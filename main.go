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

type Test struct {
	Thing string
}

func getMain(w http.ResponseWriter, r * http.Request) {
	fmt.Println("Accessing main page")

	t, _ := template.ParseFiles("index.html")
	thing := Test{"Your mom"}
	t.Execute(w, thing)
}

func main() {
	http.HandleFunc("/", getMain)

	fmt.Println("Listening at 8080")

	styles := "/static/css"
	strip := http.StripPrefix(styles, http.FileServer(http.Dir(styles[1:])))
    http.Handle("/static/css/", strip)
	http.ListenAndServe(":8080", nil)
}
