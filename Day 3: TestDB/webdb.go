package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Establishing connection with the database.
	db, err := sql.Open("mysql", "root:password@/customer_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := string(strings.Split(r.URL.Path, "/")[1])
	id, _ := strconv.Atoi(a)

	// When ID == 0, does not exist.
	if id == 0 {
		stmtOut, err := db.Query(`SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID ORDER BY Customer.ID, Address.Cus_ID;`)
		if err != nil {
			panic(err.Error())
		}
		defer stmtOut.Close()
		var res []Customer
		for stmtOut.Next() {
			var c Customer
			if err := stmtOut.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.Cus_ID); err != nil {
				log.Fatal(err)
			}
			res = append(res, c)
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err.Error())
		}

	} else {
		// Else when ID is present/ not found in DB.
		stmtOut, err := db.Query(fmt.Sprintf(`SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID WHERE Customer.ID = %v ORDER BY Customer.ID, Address.Cus_ID; `, id))
		if err != nil {
			panic(err.Error())
		}
		defer stmtOut.Close()

		var res []Customer
		for stmtOut.Next() {
			var c Customer
			if err := stmtOut.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.Cus_ID); err != nil {
				log.Fatal(err)
			}
			res = append(res, c)
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err.Error())
		}
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8070", nil))
}
