package chain

import (
	"fmt"
	"sync"
)

// The individual Links
type Link[T any] struct {
	Item  T
	next  *Link[T]
	front *Link[T]
}

// We keep track of the pointer here
type Chain[T any] struct {
	tail   *Link[T]
	head   *Link[T]
	Length int
	mu     *sync.RWMutex
}

// You can initialize now Chain struct using this
func New[T any]() *Chain[T] {
	Chain := &Chain[T]{}
	Chain.Length = 0
	return Chain
}

// Getting the item from the first Link
func (Chain Chain[T]) First() (T, error) {
	if Chain.head == nil {
		var fall T
		return fall, fmt.Errorf("chain is empty")
	}
	return Chain.head.Item, nil
}

// Getting the item from the last Link
func (Chain Chain[T]) Last() (T, error) {
	if Chain.tail == nil {
		var fall T
		return fall, fmt.Errorf("chain is empty")
	}
	return Chain.tail.Item, nil
}

// Push is adding element to the end
func (Chain *Chain[T]) Push(item T) *Chain[T] {
	new := &Link[T]{
		Item: item,
	}
	if Chain.tail == nil {
		Chain.head = new
		Chain.tail = new
	} else {
		Chain.tail.next = new
		new.front = Chain.tail
		Chain.tail = new
	}
	Chain.Length++
	return Chain
}

// front is appending element to the start
func (Chain *Chain[T]) Front(item T) *Chain[T] {
	new := &Link[T]{
		Item: item,
	}
	if Chain.head == nil {
		Chain.head = new
		Chain.tail = new
	} else {
		Chain.head.front = new
		new.next = Chain.head
		Chain.head = new
	}
	Chain.Length++
	return Chain
}

// This a way to interact with the chain, recursively from head to tail
// If your function return false it will stop
func (Chain *Chain[T]) Iter(process func(*Link[T]) bool) {
	if Chain.head == nil {
		return
	}
	var recur func(*Link[T])
	recur = func(Link *Link[T]) {
		if Link.next == nil {
			process(Link)
			return
		}

		if !process(Link) {
			return
		}
		recur(Link.next)
	}
	recur(Chain.head)
}

// This a way to interact with the chain, recursively from tail to head
// If your function return false it will stop
func (Chain *Chain[T]) Reti(process func(*Link[T]) bool) {
	if Chain.tail == nil {
		return
	}
	var recur func(*Link[T])
	recur = func(Link *Link[T]) {
		if Link.front == nil {
			process(Link)
			return
		}
		if !process(Link) {
			return
		}
		recur(Link.front)
	}
	recur(Chain.tail)
}

// Removing the last element in the Chain
func (Chain *Chain[T]) Pop() {
	if Chain.tail == nil {
		return
	}
	Chain.tail = Chain.tail.front
	Chain.tail.next = nil
}

// Removing the first element in the Chain
func (Chain *Chain[T]) Cut() {
	if Chain.head == nil {
		return
	}
	Chain.head = Chain.head.next
	Chain.head.front = nil
}

// If you have two chain, you chain them
func (Chain *Chain[T]) Merge(exChain *Chain[T]) *Chain[T] {
	if exChain.head != nil {
		exChain.head.front = Chain.tail
		Chain.tail.next = exChain.head
		Chain.tail = exChain.tail
		Chain.Length += exChain.Length
	}
	return Chain
}

// Lock for concurrency reasons, use whatever you like
func (Chain *Chain[T]) Lock() *Chain[T] {
	Chain.mu.Lock()
	return Chain
}

// Lock for concurrency reasons, use whatever you like
func (Chain *Chain[T]) Unlock() *Chain[T] {
	Chain.mu.Unlock()
	return Chain
}

// Lock for concurrency reasons, use whatever you like
func (Chain *Chain[T]) RLock() *Chain[T] {
	Chain.mu.RLock()
	return Chain
}

// Lock for concurrency reasons, use whatever you like
func (Chain *Chain[T]) RUnlock() *Chain[T] {
	Chain.mu.RUnlock()
	return Chain
}
