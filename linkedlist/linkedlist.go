package linkedlist

import "errors"

type ILinkedList interface {
	Push(data int)
	Pop() (int, error)
	Len() int
}

type node struct {
	data int
	next *node
}

type linkedlist struct {
	node *node // node.next -> node, .next -> node, .next -> nil
	len  int
}

func (ll *linkedlist) Push(data int) {
	defer func() {
		ll.len++
	}()

	nextNode := &node{
		data: data,
		next: nil,
	}

	if ll.node == nil {
		ll.node = nextNode
	} else {
		currNode := ll.node

		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = nextNode
	}
}

func (ll *linkedlist) Pop() (int, error) {
	if ll.len == 0 {
		return 0, errors.New("linked list is empty")
	}

	defer func() {
		ll.len--
	}()

	if ll.len == 1 {
		data := ll.node.data
		ll.node = nil
		return data, nil
	}

	currNode := ll.node
	for currNode.next.next != nil {
		currNode = currNode.next
	}

	data := currNode.next.data
	currNode.next = nil

	return data, nil
}

func (ll *linkedlist) Len() int {
	return ll.len
}

func New() ILinkedList {
	return &linkedlist{}
}
