package main

import (
	"fmt"
	"sync"

	"github.com/arturyumaev/linkedlist/linkedlist"
)

func main() {
	ll := linkedlist.New[int]()
	var wg sync.WaitGroup

	wg.Add(150)
	for i := 0; i < 100; i++ {
		i := i

		go func() {
			defer wg.Done()

			fmt.Println("pushing")
			ll.Push(i)
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			fmt.Println("popping")
			ll.Pop()
		}()
	}

	wg.Wait()
	fmt.Println(ll.Len())
}
