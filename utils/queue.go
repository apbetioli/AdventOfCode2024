package utils

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{value, nil}

	if q.Length == 0 {
		q.head = node
		q.tail = node
		q.Length = 1
		return
	}

	q.tail.next = node
	q.tail = node
	q.Length++
}

func (q *Queue[T]) Dequeue() T {
	if q.Length == 0 {
		panic("Empty queue")
	}

	value := q.head.value
	q.head = q.head.next
	q.Length--

	if q.Length == 0 {
		q.head = nil
		q.tail = nil
	}

	return value
}
