package iplinkedlist

import (
	"errors"
	"sync"
)

type IPNode struct {
	Prev   *IPNode
	Next   *IPNode
	IpAddr string
}

type IPLinkedList struct {
	mu sync.Mutex
	OfflineTable map[string]*IPNode
	Table map[string]*IPNode
	First *IPNode
	Last  *IPNode
	Size  int
}

func NewLinkedList(ipAddresses []string) *IPLinkedList {
	size := len(ipAddresses)
	if size == 0 {
		panic("No ip addresses is given for starting this load balancer")
	}
	currNode := &IPNode{Prev: nil, Next: nil}
	lList := IPLinkedList{First: currNode, Last: nil, Size: size, mu: sync.Mutex{}, OfflineTable: make(map[string]*IPNode)}
	ipTable := make(map[string]*IPNode)

	for i := range ipAddresses {
		currNode.IpAddr = ipAddresses[i]
		ipTable[ipAddresses[i]] = currNode
		if i <= size-2 {
			nextNode := &IPNode{Prev: currNode, Next: nil}
			currNode.Next = nextNode
			currNode = nextNode
		}
	}

	lList.Last = currNode
	lList.Table = ipTable
	return &lList
}

func (llist *IPLinkedList) GetAddr() (string, error) {
	defer llist.mu.Unlock()
	llist.mu.Lock()
	if llist.First == nil {
		return "", errors.New("all servers are not available at the moment")
	}
	
	curr := llist.First
	if curr.Next == nil {
		return curr.IpAddr, nil
	}
	addr := curr.IpAddr
	nextNode := curr.Next

	nextNode.Prev = nil
	llist.First = nextNode
	curr.Next = nil

	curr.Prev = llist.Last
	llist.Last.Next = curr

	llist.Last = curr
	return addr, nil
}

func (llist *IPLinkedList) AddToOfflineTable(node *IPNode) {
	llist.OfflineTable[node.IpAddr] = node
}
func (llist *IPLinkedList) RemoveFromOfflineTable(ipAddr string){
	delete(llist.OfflineTable, ipAddr)
}

func (llist *IPLinkedList) AddAddr(node *IPNode){
	defer llist.mu.Unlock()
	llist.mu.Lock()
	//if llist is empty?
	if llist.First == nil {
		llist.First = node
	}else{
		llist.Last.Next = node
		node.Prev = llist.Last
	}
	
	llist.Last = node
	llist.Table[node.IpAddr] = node
}

func (llist *IPLinkedList) RemoveAddr(addr string) (*IPNode, error) {
	defer llist.mu.Unlock()
	llist.mu.Lock()

	node, ok := llist.Table[addr]
	if !ok {
		return nil, errors.New("no ip address is found in this pool")
	}
	if node == llist.First && node == llist.Last {
		llist.First = nil
		llist.Last = nil
	}else {
		switch node {
	case llist.First:
		next := llist.First.Next
		if next != nil {
			next.Prev = nil
		}
		
		llist.First = next
	case llist.Last:
		prev := node.Prev
		prev.Next = nil
		llist.Last = prev
	default:
		prev := node.Prev
		next := node.Next

		prev.Next = next
		next.Prev = prev

	}
	}
	
	delete(llist.Table, addr)
	node.Next = nil
	node.Prev = nil

	return node, nil
}