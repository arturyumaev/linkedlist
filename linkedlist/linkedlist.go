package linkedlist

import "errors"

type ILinkedList interface {
	Push(data int)
	Pop() (int, error)
	PopNth(n int) (int, error)
	Len() int
}

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node // node.next -> node.next -> node.next -> nil
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

	if ll.head == nil {
		ll.head = nextNode
	} else {
		currNode := ll.head

		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = nextNode
	}
}

func (ll *linkedlist) Pop() (int, error) {
	return ll.PopNth(ll.len - 1)
}

func (ll *linkedlist) PopNth(n int) (int, error) {
	if ll.len == 0 {
		return 0, errors.New("linked list is empty")
	}

	if n < 0 || n >= ll.len {
		return 0, errors.New("n is outside of range")
	}

	if ll.len == 1 {
		data := ll.head.data
		ll.head = nil
		ll.len--
		return data, nil
	}

	currNode := ll.head
	for c := 0; c < n-1; c++ {
		currNode = currNode.next
	}
	data := currNode.next.data
	currNode.next = currNode.next.next
	ll.len--

	return data, nil
}

func (ll *linkedlist) Len() int {
	return ll.len
}

func New() ILinkedList {
	return &linkedlist{}
}
