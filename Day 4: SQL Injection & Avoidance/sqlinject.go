package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID   int
	Name string
	DOB  string
	Lang string
}

func handler2(w http.ResponseWriter, r *http.Request) {
	// Establishing connection with the database.
	db, err := sql.Open("mysql", "root:password@/customer_service?multiStatements=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := strings.Split(r.URL.Path, "/")[1]
	stmt1 := strings.Split(id, ";")[0]
	//stmt2 := strings.Split(id, ";")[1]
	//fmt.Println("id is ", id)
	//var ids []interface{}
	// When ID == 0, does not exist.
	query := `SELECT * FROM DropIt`
	// If id is given.
	if len(stmt1) != 0 {
		// Using placeholders for SQL injection avoidance, placeholder = ?
		//query = `SELECT * FROM DropIt WHERE ID = ?`
		query = "SELECT * FROM DropIt WHERE ID = " + id
	}
	rows, err := db.Query(query) //ids...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var res []Customer
	//for rows.Next() {
	//	var c Customer
	//	if err := rows.Scan(&c.ID, &c.Name, &c.Lang); err != nil {
	//		log.Fatal(err)
	//	}
	//	res = append(res, c)
	//}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe(":10070", nil))
	// To free the server use:
	// sudo kill -9 `sudo lsof -t -i:10070`
}
