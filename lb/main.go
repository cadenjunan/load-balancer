package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"test-loadbalancer/iplinkedlist"
)



type WebRequest struct {
	client *http.Client
	ipLinkedList *iplinkedlist.IPLinkedList
}

func ErrorHandler(w *http.ResponseWriter, err error){
	log.Println(err.Error())
	http.Error(*w, err.Error(), http.StatusInternalServerError)
}

func (wr *WebRequest) Reroute(w http.ResponseWriter, r *http.Request){
	method := r.Method
	body:= r.Body
	
	defer body.Close()
	uri, err := wr.ipLinkedList.GetAddr()
	if err != nil {
		ErrorHandler(&w,err)
		return
	}
	req, err := http.NewRequest(method, uri + r.RequestURI, body)
	if err != nil {
		ErrorHandler(&w,err)
		return
	}
	for h := range r.Header {
		req.Header.Add(h, r.Header.Get(h))
	}
	
	if resp, err := wr.client.Do(req); err != nil {
		ErrorHandler(&w,err)
	}else{
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ErrorHandler(&w,err)
		}else{
			fmt.Fprintf(w, "%s", string(body))
		}
		
	}
}

func main() {
	wr := WebRequest{client: &http.Client{}, ipLinkedList: iplinkedlist.NewLinkedList([]string{"http://localhost:8000","http://localhost:8001", "http://localhost:8002"})}
	log.Println("Load balancer runs at localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(wr.Reroute)))
}