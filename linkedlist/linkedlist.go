package linkedlist

import "errors"

type ILinkedList[T any] interface {
	Push(data T)
	Pop() (T, error)
	PopNth(n int) (T, error)
	Len() int
}

type node[T any] struct {
	data T
	next *node[T]
}

type linkedlist[T any] struct {
	head *node[T] // node.next -> node.next -> node.next -> nil
	len  int
}

func (ll *linkedlist[T]) Push(data T) {
	defer func() {
		ll.len++
	}()

	nextNode := &node[T]{
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

func (ll *linkedlist[T]) Pop() (T, error) {
	return ll.PopNth(ll.len - 1)
}

func (ll *linkedlist[T]) PopNth(n int) (T, error) {
	var defaultValue T

	if ll.len == 0 {
		return defaultValue, errors.New("linked list is empty")
	}

	if n < 0 || n >= ll.len {
		return defaultValue, errors.New("n is outside of range")
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

func (ll *linkedlist[T]) Len() int {
	return ll.len
}

func New[T any]() ILinkedList[T] {
	return &linkedlist[T]{}
}
