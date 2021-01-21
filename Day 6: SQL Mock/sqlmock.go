package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	Id       int
	Username string
}

func GetCustomer(db *sql.DB) (err error) {
	query := `insert into Customers values`
	_, err = db.Query(query)

	if err != nil {
		fmt.Println("Error: Query cant take place, SQL mock connection issue", err)
	}

	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM Customers;")

	if err != nil {
		fmt.Println("Error is ", err)
	}

	var res []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Username); err != nil {
			log.Fatal(err)
		}
		res = append(res, c)
	}
	return err
}

//func main() {
//	db, err := sql.Open("mysql", "root:nimda@/test")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	err = cancelOrder(1, db)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
