package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := call()
	fmt.Fprintf(w, "Router1: %s \nRouter2: %s error: %s", r.URL.Path[1:], res, err)
}

func call() (string, error) {
	req, err := http.NewRequest("GET", "http://linkerd_router2", nil)
	if err != nil {
		return "failed", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "failed", err
	}
	defer resp.Body.Close()

	var res string
	fmt.Println("decode...")
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Println(err)
	}
	fmt.Println("end...")
	fmt.Println(res)
	return res, nil
}

func main() {
	fmt.Println("server started on port :8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
