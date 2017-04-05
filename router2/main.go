package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Router2: received %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("server started on port :8020...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8020", nil)
}
