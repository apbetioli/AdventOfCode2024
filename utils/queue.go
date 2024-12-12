package utils

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	Length int
	head   *Node
	tail   *Node
}

func (q *Queue) Enqueue(value int) {
	node := &Node{value, nil}

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

func (q *Queue) Dequeue() int {
	if q.Length == 0 {
		return -1
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
