package gr

type node struct {
	element interface{}
	next    *node
}

type queue struct {
	head *node
	tail *node
	size int
}

func (q *queue) enqueue(element interface{}) {

	n := new(node)
	n.element = element

	if q.size == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
	q.size++
}

func (q *queue) dequeue() interface{} {

	if q.size == 0 {
		return nil
	} else {
		n := q.head
		r := q.head.element
		q.size--

		if q.size == 0 {
			q.head = nil
			q.tail = nil
		} else {
			q.head = q.head.next
		}

		n.element = nil
		n.next = nil
		n = nil

		return r
	}
}

func (q *queue) Size() int {
	return q.size
}
