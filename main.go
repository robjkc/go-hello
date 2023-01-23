package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
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

	initPostgres()
	initFarad()

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

func initPostgres() {
	fmt.Println("preparing to lookup postgres.local.")
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
}

func initFarad() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d",
		"farad.local", "rjohnsonscw", "t1oyota2", "farad_production", 1433)
	fmt.Printf(" connString:%s\n", connString)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select id, name from control_systems where id = 24")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var controlSystemID int64
	var controlSystemName string
	err = row.Scan(&controlSystemID, &controlSystemName)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("controlSystemID:%d\n", controlSystemID)
	fmt.Printf("controlSystemName:%s\n", controlSystemName)

	fmt.Printf("bye\n")
}
