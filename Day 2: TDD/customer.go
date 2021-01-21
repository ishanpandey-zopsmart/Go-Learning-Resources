package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	Age     int
	Address string
}

// Test directly calls handler func to test for outputs.
func handler(w http.ResponseWriter, r *http.Request) {
	// reads bytes from input output.
	body, _ := ioutil.ReadAll(r.Body)
	// read the bytes, after converting them into type interface/ struct here, using unmarshal.
	// Create an output var customer
	var c Customer
	json.Unmarshal(body, &c)

	/*
		If you handle error, and body is empty i.e. localhost:8070/ , then body is nil so unmarshal will give error, so log.Fataf will stop the execution.
		err :=
		if err != nil {
		log.Fatalf("%s", err)
		}
	*/

	//fmt.Println(c)

	if c.Age <= 18 {
		io.WriteString(w, "not eligible")
	} else {
		// Can use \ (backslash) to ignore " in fmt.sprintf for formatting.
		// 1. fmt.Printf
		//io.WriteString(w, fmt.Sprintf("{\"Name\" : \"%s\", \"Age\" : %v, \"Address\" : \"%s\"}", c.Name, c.Age, c.Address))

		// 2. using simple body.
		//io.WriteString(w, string(body))

		// 3. using JSON Marshal : (only merges all key value pairs, without spaces and just, in between JSON)
		jsonData, err := json.Marshal(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error occurred!")
		}
		io.WriteString(w, string(jsonData))
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	if string(r.Method) == "GET" {
		io.WriteString(w, "GET - Hello")
	} else {
		io.WriteString(w, "POST - Hello")
	}
}

// Main spawns the server.
func main() {
	//http.HandleFunc("/printrtype", handler1)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8070", nil))
}

// For code coverage:
// go test ./... -v -coverprofile coverage.txt -coverpkg=./...
// go tool cover -html coverage.txt
