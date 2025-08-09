package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func helloMessage(w http.ResponseWriter, r *http.Request){
	log.Printf("Host %s serving request at %s\n", r.Host, r.RequestURI)
	m := Message{Message: "hello world!"}
	encoder := json.NewEncoder(w)
	encoder.Encode(m)
}

func main() {
	port := flag.String("port", "8000", "Enter the port number")
	flag.Parse()
	http.HandleFunc("/hello", helloMessage)
	log.Printf("Starting server at port %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}