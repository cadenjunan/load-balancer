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

func TestOneIpNode(t *testing.T){
	ipAddresses := []string{"192.168.0.1"}
	llist := NewLinkedList(ipAddresses)
	for range 5 {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != "192.168.0.1" {
			t.Error("address is not 192.168.0.1")
		}
	}
	node, _ := llist.RemoveAddr("192.168.0.1")
	if llist.First != nil || llist.Last != nil {
		t.Error("fail to clear first and last reference when llist is empty")
	}
	_, err := llist.GetAddr()
	if err == nil {
		t.Error("fail to throw error when llist is empty")
	}
	llist.AddAddr(node)
	llist.AddAddr(&IPNode{IpAddr: "192.168.0.2"})
	modifiedIpAddresses := []string{"192.168.0.1", "192.168.0.2"}

	for i:= range 2 {
		addr, _ := llist.GetAddr()
		if addr != modifiedIpAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedIpAddresses[i], addr)
		}
	}
	for i:= range 2 {
		addr, _ := llist.GetAddr()
		if addr != modifiedIpAddresses[i] {
			t.Errorf("Wanted %s but got %s", modifiedIpAddresses[i], addr)
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
func TestRemoveAddIpAddr(t *testing.T) {
	ipAddresses := []string{ "192.168.0.1","192.168.0.2", "192.168.0.3", "192.168.0.4" }
	llist := NewLinkedList(ipAddresses)
	node, _ :=llist.RemoveAddr("192.168.0.2")
	_, ok := llist.Table[node.IpAddr]
	if ok {
		t.Errorf("ip address should not be found in the table after its removal")
	}
	llist.AddAddr(node)
	_, ok = llist.Table[node.IpAddr]
	if !ok {
		t.Errorf("ip address should be found in the table after adding back into the table")
	}

	modifiedAddresses := []string{"192.168.0.1",  "192.168.0.3","192.168.0.4","192.168.0.2"}
	for _, addr := range modifiedAddresses {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != addr {
			t.Errorf("Wanted %s but got %s",addr, ipAddr)
		}
	}
}

func TestRemoveAddOneAddr(t *testing.T){
	ipAddresses := []string{"192.168.0.1"}
	llist := NewLinkedList(ipAddresses)
	node, _ :=llist.RemoveAddr("192.168.0.1")
	if len(llist.Table) != 0 {
		t.Error("fail to delete entry 192.168.0.1 from the ip table")
	}
	if llist.First != nil || llist.Last != nil {
		t.Errorf("fail to clear off first and last node %s", node.IpAddr)
	}
	llist.AddAddr(node)
	if len(llist.Table) != 1 {
		t.Error("fail to add entry 192.168.0.1 to the ip table")
	}
	
	for range 3 {
		ipAddr, _ := llist.GetAddr()
		if ipAddr != "192.168.0.1" {
			t.Error("fail to get the same address of 192.168.0.1")
		}
	}
}