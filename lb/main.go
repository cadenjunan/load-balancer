package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type UriRegistry struct {
	mu sync.Mutex
	addresses []string
	currIndex int
	endIndex int
}

func NewRegistry(addresses []string) *UriRegistry {
	if len(addresses) == 0 {
		panic("List of ip addresses is empty.")
	}
	return &UriRegistry{
		mu: sync.Mutex{},
		addresses: addresses,
		currIndex: 0,
		endIndex: len(addresses)-1,
	}
}
func (r *UriRegistry) GetIpAddress() string {
	defer r.mu.Unlock()
	r.mu.Lock()
	if r.currIndex > r.endIndex {
		r.currIndex = 0
	}
	index := r.currIndex
	r.currIndex++
	return r.addresses[index]
}

type WebRequest struct {
	client *http.Client
	registry *UriRegistry
}

func ErrorHandler(w *http.ResponseWriter, err error){
	log.Println(err.Error())
	http.Error(*w, err.Error(), http.StatusInternalServerError)
}

func (wr *WebRequest) Reroute(w http.ResponseWriter, r *http.Request){
	method := r.Method
	body:= r.Body
	
	defer body.Close()
	uri := wr.registry.GetIpAddress()
	req, err := http.NewRequest(method, uri + r.RequestURI, body)
	if err != nil {
		ErrorHandler(&w,err)
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
	wr := WebRequest{client: &http.Client{}, registry: NewRegistry([]string{"http://localhost:8000","http://localhost:8001", "http://localhost:8002"})}
	log.Println("Load balancer runs at localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(wr.Reroute)))
}