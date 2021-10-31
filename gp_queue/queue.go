package gp_queue

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	n := &Node{}
	return &Queue{
		head: n,
		tail: n,
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}
	v := q.head.next.value
	q.head = q.head.next
	return v
}

func (q *Queue) Enqueue(v interface{}) {
	n := &Node{
		value: v,
		next:  nil,
	}
	q.tail.next = n
	q.tail = n
}
