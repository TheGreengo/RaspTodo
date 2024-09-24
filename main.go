package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r * http.Request) {
	fmt.Println("access")
	io.WriteString(w, "this is working")
}

func main() {
	http.HandleFunc("/", getRoot)
	fmt.Println("Listening at 8080")
	http.ListenAndServe(":8080", nil)
}
