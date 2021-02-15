package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/sendData", sendData)
	http.ListenAndServe(":8090", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func sendData(w http.ResponseWriter, req *http.Request) {
	send()
	fmt.Fprintf(w, "hello\n")
}