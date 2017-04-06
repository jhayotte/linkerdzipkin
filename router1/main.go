package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := call(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Router1: %s \nRouter2: error: %s", r.URL.Path[1:], err)
		return
	}
	fmt.Fprintf(w, "Router1: root %s \n%s", r.URL.Path[1:], res)
}

func call(path string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://linkerd_router1:8090/%s", path))
	if err != nil {
		return "failed", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	fmt.Println("server started on port :8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
