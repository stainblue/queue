package queue

type Queue struct {
	head *queueItem
	tail *queueItem
	cnt  int
}

type queueItem struct {
	item *interface{}
	next *queueItem
}

func New() *Queue {
	return new(Queue)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return q.cnt
}

func (q *Queue) Enqueue(value interface{}) {
	if q.cnt == 0 {
		q.tail = &queueItem{
			item: &value,
			next: nil,
		}
		q.head = q.tail
	} else {
		q.tail.next = &queueItem{
			item: &value,
			next: nil,
		}
		q.tail = q.tail.next
	}
	q.cnt++
}

func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}

	output := q.head.item
	q.head = q.head.next
	q.cnt--

	if q.cnt == 0 {
		q.tail = nil
	}

	return *output
}
