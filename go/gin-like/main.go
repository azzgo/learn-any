package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHanlder)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func helloHanlder(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "URL.Path = %q\n", req.URL.Path)
}
