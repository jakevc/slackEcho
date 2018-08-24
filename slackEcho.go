package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// handle the http response and request
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dat := r.Form
	echoText := strings.Join(dat["text"], "")

	fmt.Fprintf(w, echoText)
}

func main() {
	// handle request and echo
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5038", nil))
}
