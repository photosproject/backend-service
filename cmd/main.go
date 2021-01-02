package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/hello", responseMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func responseMessage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"Hello World!")
}


