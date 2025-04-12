package double

import (
	"fmt"
)

const (
	TAIL = -1
	HEAD = 0
)

// The individual Links
type Link[T any] struct {
	Item  T
	Next  *Link[T]
	Front *Link[T]
}

func (link *Link[T]) Delete() {
	link.Detach()
	link = nil
}

// Removing oneself from the queue
// whilst maintaining it's structure
func (link *Link[T]) Detach() {
	link.Front.Next = link.Next
	link.Next.Front = link.Front
}

// We keep track of the pointer here
type Chain[T any] struct {
	tail   *Link[T]
	head   *Link[T]
	Mark   *Link[T]
	Length int
}

// You can initialize new Chain struct using this
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

// Push is adding element to the back
func (Chain *Chain[T]) Push(item T) *Chain[T] {
	new := &Link[T]{
		Item: item,
	}
	if Chain.tail == nil {
		Chain.head = new
		Chain.tail = new
	} else {
		Chain.tail.Next = new
		new.Front = Chain.tail
		Chain.tail = new
	}
	Chain.Length++
	return Chain
}

// Front is appending element to well, Front
func (Chain *Chain[T]) Front(item T) *Chain[T] {
	new := &Link[T]{
		Item: item,
	}
	if Chain.head == nil {
		Chain.head = new
		Chain.tail = new
	} else {
		Chain.head.Front = new
		new.Next = Chain.head
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
		if Link.Next == nil {
			process(Link)
			return
		}

		if !process(Link) {
			return
		}
		recur(Link.Next)
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
		if Link.Front == nil {
			process(Link)
			return
		}
		if !process(Link) {
			return
		}
		recur(Link.Front)
	}
	recur(Chain.tail)
}

// Removing the last element in the Chain
func (Chain *Chain[T]) Pop() {
	if Chain.tail == nil {
		return
	}
	Chain.tail = Chain.tail.Front
	Chain.tail.Next = nil

	Chain.Length--
}

// Removing the first element in the Chain
func (Chain *Chain[T]) Cut() {
	if Chain.head == nil {
		return
	}
	Chain.head = Chain.head.Next
	Chain.head.Front = nil
	Chain.Length--
}

// If you have two chain, you can chain them
// DO NOT TRY TO MERGE WITH THE SAME CHAIN TWICE
// Ie doing this:
// Chain.Merge(other)
// Chain.Merge(other)
// IS A BAD IDEA!
func (Chain *Chain[T]) Merge(exChain *Chain[T]) *Chain[T] {
	exChain.head.Front = Chain.tail
	Chain.tail.Next = exChain.head
	Chain.tail = exChain.tail
	Chain.Length += exChain.Length
	return Chain
}

func (Chain *Chain[T]) Attach(link *Link[T]) {
	Chain.tail.Next = link
	link.Front = Chain.tail
	Chain.tail = link
	Chain.Length++
}

// Retrieve the head pointer beware if it's nil
func (Chain *Chain[T]) Head() *Link[T] {
	return Chain.head
}

// Retrieve the tail pointer beware if it's nil
func (Chain *Chain[T]) Tail() *Link[T] {
	return Chain.tail
}

// You mark a pointer and then retrive it's value
// default to the Tail
func (Chain *Chain[T]) Point(i int) *Link[T] {
	switch i {
	case HEAD:
		Chain.Mark = Chain.head
	case TAIL:
		Chain.Mark = Chain.tail
	default:
		count := 0
		Chain.Iter(func(chain *Link[T]) bool {
			Chain.Mark = chain
			count++
			return count != i
		})
	}
	return Chain.Mark
}

// You move the mark Up towards the head one level
// Then retrieve it's pointer
// You need to to
func (Chain *Chain[T]) Up() *Link[T] {
	if Chain.Mark != nil {
		Chain.Point(0)
	}
	Chain.Mark = Chain.Mark.Front
	return Chain.Mark
}

// You move the mark Down towards the tail one level
// Then retrieve it's pointer
func (Chain *Chain[T]) Down() *Link[T] {
	if Chain.Mark != nil {
		Chain.Point(0)
	}
	Chain.Mark = Chain.Mark.Next
	return Chain.Mark
}

func (Chain *Chain[T]) Clear() {
	var recur func(*Link[T])
	recur = func(Link *Link[T]) {
		if Link.Next == nil {
			Link = nil
			return
		}
		tmp := Link.Next
		Link = nil
		recur(tmp)
	}
	recur(Chain.head)
}
