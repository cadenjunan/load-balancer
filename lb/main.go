package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type UriRegistry struct {
	uri string
}

type WebRequest struct {
	client *http.Client
	registry UriRegistry
}



func (wr *WebRequest) Reroute(w http.ResponseWriter, r *http.Request){
	method := r.Method
	body:= r.Body
	
	
	defer body.Close()
	req, err := http.NewRequest(method, wr.registry.uri + r.RequestURI, body)
	if err != nil {
		log.Panic(err.Error())
	}
	for h := range r.Header {
		req.Header.Add(h, r.Header.Get(h))
	}
	
	if resp, err := wr.client.Do(req); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}else{
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err.Error())
		}
		fmt.Fprintf(w, "%s", string(body))
	}
}

func main() {
	wr := WebRequest{client: &http.Client{}, registry: UriRegistry{uri: "http://localhost:8000"}}
	log.Println("Load balancer runs at localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(wr.Reroute)))
}