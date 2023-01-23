package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net"
	"net/http"
)

const (
	host     = "postgres.local"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/go", HelloGo)
	http.HandleFunc("/control-system", GetControlSystem)

	fmt.Println("preparing to lookup postgres-svc.")
	ips, err := net.LookupIP("postgres.local" +
		"")
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("Prepare to open!")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping!")
	} else {
		fmt.Println("Established a successful connection!")
	}

	sqlStatement := `SELECT id, name FROM test WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, 1)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		fmt.Println("No data!")
		fmt.Println(err)
	}

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

	fmt.Fprintf(w, "{\"msg\":\"Hello from Go5!\"}")
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
