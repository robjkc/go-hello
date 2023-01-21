package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/go", HelloGo)
	http.HandleFunc("/control-system", GetControlSystem)

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

	fmt.Fprintf(w, "{\"msg\":\"Hello from Go1!\"}")
	fmt.Println("Go endpoint hit!")
}

func GetControlSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")

	fmt.Println("Control System endpoint hit!")

	url := "https://apid.securecomwireless.com/v2/control_systems/12345?auth_token=wR8aZxxZAuhyhtaH2iyb"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, string(body))
}
