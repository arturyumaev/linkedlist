package linkedlist

import (
	"errors"
	"sync"
)

type ILinkedList[T any] interface {
	Push(data T)
	Pop() (T, error)
	Len() int
}

type node[T any] struct {
	data T
	next *node[T]
}

type linkedlist[T any] struct {
	head *node[T]
	len  int
	mu   *sync.RWMutex
}

func (ll *linkedlist[T]) Push(data T) {
	nextNode := &node[T]{
		data: data,
		next: nil,
	}

	ll.mu.Lock()
	defer ll.mu.Unlock()

	if ll.head == nil {
		ll.head = nextNode
	} else {
		currNode := ll.head

		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = nextNode
	}

	ll.len++
}

func (ll *linkedlist[T]) Pop() (T, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	var defaultValue T
	currentLen := ll.len

	if currentLen == 0 {
		return defaultValue, errors.New("linked list is empty")
	}

	if currentLen == 1 {
		data := ll.head.data
		ll.head = nil
		ll.len--
		return data, nil
	}

	currNode := ll.head
	for currNode.next.next != nil {
		currNode = currNode.next
	}
	data := currNode.next.data
	currNode.next = nil
	ll.len--

	return data, nil
}

func (ll *linkedlist[T]) Len() int {
	ll.mu.RLock()
	len := ll.len
	ll.mu.RUnlock()
	return len
}

func New[T any]() ILinkedList[T] {
	return &linkedlist[T]{
		mu: &sync.RWMutex{},
	}
}
