package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))

}

func homeHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	for k, v := range req.Header {

		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)

	}

}