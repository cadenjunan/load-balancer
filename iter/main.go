package main
import (
	"fmt"
)

type Registry struct {
	ipAddress []string
	currIndex int
	endIndex int
}
func NewRegistry(ipAddress []string) *Registry {
	return &Registry{
		currIndex: 0,
		endIndex: len(ipAddress)-1,
		ipAddress: ipAddress,
	}
}
func (r *Registry) GetIpAddress() string {
	if r.currIndex > r.endIndex {
		r.currIndex = 0
	}
	index := r.currIndex
	r.currIndex++
	return r.ipAddress[index]
}


func main(){
	r := NewRegistry([]string{"192.168.0.11:8080", "192.168.11.1:8001","192.168.0.125:8001"})
	fmt.Println(r.GetIpAddress())
	fmt.Println(r.GetIpAddress())
	fmt.Println(r.GetIpAddress())
	fmt.Println(r.GetIpAddress())
	fmt.Println(r.GetIpAddress())
}