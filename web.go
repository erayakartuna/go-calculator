package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9090",nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are in : "+ r.URL.Path)
}
