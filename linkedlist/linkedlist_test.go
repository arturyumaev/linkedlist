package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Pop(t *testing.T) {
	ll := New[int]()

	v, err := ll.Pop()
	assert.Equal(t, v, 0)
	assert.ErrorContains(t, err, "empty")

	v, err = ll.Pop()
	assert.Equal(t, v, 0)
	assert.ErrorContains(t, err, "empty")
}

func TestLinkedList_Push(t *testing.T) {
	ll := New[int]()
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	assert.Equal(t, 3, ll.Len())

	v, err := ll.Pop()
	assert.Equal(t, 3, v)
	assert.NoError(t, err)
	assert.Equal(t, 2, ll.Len())

	v, err = ll.Pop()
	assert.Equal(t, 2, v)
	assert.NoError(t, err)
	assert.Equal(t, 1, ll.Len())

	v, err = ll.Pop()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
	assert.Equal(t, 0, ll.Len())

	v, err = ll.Pop()
	assert.Equal(t, 0, v)
	assert.ErrorContains(t, err, "empty")
	assert.Equal(t, 0, ll.Len())
}
