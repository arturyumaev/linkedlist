package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("edge cases", func(t *testing.T) {
		ll := New()

		v, err := ll.Pop()
		assert.Equal(t, v, 0)
		assert.ErrorContains(t, err, "empty")

		v, err = ll.Pop()
		assert.Equal(t, v, 0)
		assert.ErrorContains(t, err, "empty")
	})

	t.Run("test success", func(t *testing.T) {
		ll := New()
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
	})
}
