package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("Listening on 8082")
	http.ListenAndServe(":8082", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There5, %s!", r.URL.Path[1:])
	fmt.Println("Hello endpoint hit!")
}
