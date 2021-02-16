package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/sendData", sendData)
	http.HandleFunc("/sendData2", sendData2)
	http.HandleFunc("/getData", getData)
	http.ListenAndServe(":8090", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {


	fmt.Fprintf(w, "hello\n")
}

func sendData(w http.ResponseWriter, req *http.Request) {
	var index, data string
	for k, v := range req.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
		if k == "index" {
			index = v[0]
		}

		if k == "data" {
			data = v[0]
		}
	}
	send(index, data)
	fmt.Fprintf(w, "123")
}

func sendData2(w http.ResponseWriter, req *http.Request) {
	var index, data string
	for k, v := range req.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
		if k == "index" {
			index = v[0]
		}

		if k == "data" {
			data = v[0]
		}
	}
	send2(index, data)
	fmt.Fprintf(w, "123")
}

func getData(w http.ResponseWriter, req *http.Request) {
	var data string
	for k, v := range req.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
		if k == "index" {
			data = get(v[0])
		}
	}
	fmt.Fprintf(w, data)
}