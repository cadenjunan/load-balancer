package iplinkedlist

import (
	"testing"
)

func TestRoundRobin(t *testing.T) {
	ipAddresses := []string{"192.168.0.1", "192.16.0.2", "192.168.0.3"}
	llist := NewLinkedList(ipAddresses)
	size := len(ipAddresses)
	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != ipAddresses[i] {
			t.Errorf("Wanted %s but got %s", ipAddresses[i], ipAddr)
		}
	}
	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != ipAddresses[i] {
			t.Errorf("Wanted %s but got %s", ipAddresses[i], ipAddr)
		}
	}
	
}
func TestRemoveIPNode(t *testing.T){
	ipAddresses := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3",  "192.168.0.4"}
	llist := NewLinkedList(ipAddresses)
	node, _ := llist.RemoveAddr("192.168.0.3")
	
	if node.IpAddr != "192.168.0.3" {
		t.Errorf("remove %s address, but got %s", "192.168.0.3", node.IpAddr)

	}else if node.Prev != nil && node.Next != nil {
		t.Errorf("Both the prev and next link is not cleared in the removed node")
	}
	//check if removed node is not the ip table
	if _, ok := llist.Table["192.168.0.3"]; ok {
		t.Errorf("192.168.0.3 is not removed from the ip table")
	}
	modifiedAddresses := []string{"192.168.0.1", "192.168.0.2",  "192.168.0.4"}
	size := len(modifiedAddresses)

	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
}

func TestRemoveFirstIPNode(t *testing.T){
	ipAddresses := []string{ "192.168.0.1","192.168.0.2", "192.168.0.3",  "192.168.0.4"}
	llist := NewLinkedList(ipAddresses)
	node, _ := llist.RemoveAddr("192.168.0.1")
	
	if node.IpAddr != "192.168.0.1" {
		t.Errorf("remove %s address, but got %s", "192.168.0.1", node.IpAddr)

	}else if node.Prev != nil && node.Next != nil {
		t.Errorf("Both the prev and next link is not cleared in the removed node")
	}
	//check if removed node is not the ip table
	if _, ok := llist.Table["192.168.0.1"]; ok {
		t.Errorf("192.168.0.1 is not removed from the ip table")
	}
	modifiedAddresses := []string{"192.168.0.2", "192.168.0.3",  "192.168.0.4"}
	size := len(modifiedAddresses)

	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
}

func TestRemoveLastIPNode(t *testing.T){
	ipAddresses := []string{ "192.168.0.1","192.168.0.2", "192.168.0.3",  "192.168.0.4"}
	llist := NewLinkedList(ipAddresses)
	node, _ := llist.RemoveAddr("192.168.0.4")
	
	if node.IpAddr != "192.168.0.4" {
		t.Errorf("remove %s address, but got %s", "192.168.0.4", node.IpAddr)

	}else if node.Prev != nil && node.Next != nil {
		t.Errorf("Both the prev and next link is not cleared in the removed node")
	}
	//check if removed node is not the ip table
	if _, ok := llist.Table["192.168.0.4"]; ok {
		t.Errorf("192.168.0.4 is not removed from the ip table")
	}
	modifiedAddresses := []string{"192.168.0.1", "192.168.0.2",  "192.168.0.3"}
	size := len(modifiedAddresses)

	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
	for i := range size {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != modifiedAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedAddresses[i], ipAddr)
		}
	}
}

func TestEmptyIpAddressList(t *testing.T){
	defer func(){
		if r := recover(); r == nil {
			t.Errorf("linked list instantiation did not panic when pass empty list of ip address")
		}
	}()
	NewLinkedList([]string{})
}

func TestNoAvailableIPAddress(t *testing.T){
	llist := NewLinkedList([]string{"192.168.0.1"})
	llist.RemoveAddr("192.168.0.1")
	ipAddr, err := llist.GetAddr()
	if err == nil {
		t.Errorf("no error is thrown when there is no available ip address in the pool")
	}
	if ipAddr != ""{
		t.Errorf("should have return empty string, but received %s", ipAddr)
	}
}