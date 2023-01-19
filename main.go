package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/go", HelloGo)
	fmt.Println("Listening on 8082")
	http.ListenAndServe(":8082", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Println("Hello endpoint hit!")
}

func HelloGo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")

	fmt.Fprintf(w, "{\"msg\":\"Hello from Go New version!\"}")
	fmt.Println("Go endpoint hit!")
}
