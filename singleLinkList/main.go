package singleLinkList

import (
	"fmt"

	"github.com/recor-glitch/LinkList/hash"
)

type Data struct {
	hash  string
	value any
}

type Node struct {
	next *Node
	data Data
}

type LinkList struct {
	head *Node
}

func (ll *LinkList) Insert(key string, data any) (bool, error) {
	hashedKey, err := hash.GetHash(key)
	if err != nil {
		return false, err
	}
	newNode := &Node{data: Data{
		hash:  string(hashedKey[:]),
		value: data,
	}}

	if ll.head == nil {
		ll.head = newNode
		return true, nil
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return true, nil
}

func (ll *LinkList) Delete(key string) (bool, error) {
	if ll.head == nil {
		return false, fmt.Errorf("there is no head")
	}

	hashedKey, err := hash.GetHash(key)
	if err != nil {
		return false, err
	}

	if ll.head.data.hash == string(hashedKey[:]) {
		ll.head = ll.head.next
		return true, nil
	}

	current := ll.head
	for current.next != nil && current.data.hash == string(hashedKey[:]) {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
	return true, nil
}

func (ll *LinkList) Seek(key string) (any, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("no data found for key: %s", key)
	}

	hashedKey, err := hash.GetHash(key)
	if err != nil {
		return nil, err
	}

	if ll.head.data.hash == string(hashedKey[:]) {
		return ll.head.data.value, nil
	}

	current := ll.head

	for current != nil {
		if current.data.hash == string(hashedKey[:]) {
			return current.data.value, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("no data found for key: %s", key)
}

func (ll *LinkList) Display() {
	current := ll.head

	for current != nil {
		fmt.Printf("%+v\n", current.data.value)
		current = current.next
	}
	fmt.Printf("nil")

}
