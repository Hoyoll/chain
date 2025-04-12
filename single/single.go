package single

// The individual Links
type Link[T any] struct {
	Item T
	Next *Link[T]
}

// We keep track of the pointer here
type Chain[T any] struct {
	head   *Link[T]
	tail   *Link[T]
	Length int
}

func New[T any]() *Chain[T] {
	Chain := &Chain[T]{}
	Chain.Length = 0
	return Chain
}

func (Chain *Chain[T]) Push(item T) *Chain[T] {
	new := &Link[T]{
		Item: item,
	}
	if Chain.tail == nil {
		Chain.head = new
		Chain.tail = new
	} else {
		Chain.tail.Next = new
		Chain.tail = new
	}
	Chain.Length++
	return Chain
}

// Removing the last element in the Chain
func (Chain *Chain[T]) Pop() {
	if Chain.tail == nil {
		return
	}
	Chain.Iter(func(l *Link[T]) bool {
		if l.Next.Next != nil {
			return true
		}
		l.Next = nil
		Chain.tail = l
		return false
	})
	Chain.Length--
}

// Removing the first element in the Chain
func (Chain *Chain[T]) Cut() {
	if Chain.head == nil {
		return
	}
	var temp *Link[T]
	temp = Chain.head
	Chain.head = temp.Next
	temp = nil
	Chain.Length--
}

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

func (Chain *Chain[T]) Attach(link *Link[T]) {
	Chain.tail.Next = link
	Chain.tail = link
	Chain.Length++
}

func (Chain *Chain[T]) Merge(exChain *Chain[T]) *Chain[T] {
	Chain.tail.Next = exChain.head
	Chain.tail = exChain.tail
	Chain.Length += exChain.Length
	return Chain
}
