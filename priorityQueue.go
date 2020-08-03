package queue

import "container/heap"

type priorityQueue struct {
	implementation priorityQueueImplement
}

func NewPriorityQueue(compareFunc func(i, j interface{}) bool) *priorityQueue {
	pqi := priorityQueueImplement{
		items: make([]*interface{}, 0),
		compareFunc: compareFunc,
	}
	heap.Init(&pqi)

	return &priorityQueue{implementation: pqi}
}

func (p *priorityQueue) IsEmpty() bool {
	return p.Size() == 0
}

func (p *priorityQueue) Size() int {
	return p.implementation.Len()
}

func (p *priorityQueue) Push(value interface{}) {
	heap.Push(&p.implementation, value)
}

func (p *priorityQueue) Pop() interface{} {
	return heap.Pop(&p.implementation)
}

// A priorityQueueImplement implements heap.Interface.
type priorityQueueImplement struct {
	items       []*interface{}
	compareFunc func(i, j interface{}) bool
}

func (pqi priorityQueueImplement) Len() int {
	return len(pqi.items)
}

func (pqi priorityQueueImplement) Less(i, j int) bool {
	return pqi.compareFunc(*(pqi.items[i]), *(pqi.items[j]))
}

func (pqi priorityQueueImplement) Swap(i, j int) {
	pqi.items[i], pqi.items[j] = pqi.items[j], pqi.items[i]
}

func (pqi *priorityQueueImplement) Push(x interface{}) {
	pqi.items = append(pqi.items, &x)
}

func (pqi *priorityQueueImplement) Pop() interface{} {
	n := len(pqi.items)
	item := pqi.items[n-1]
	pqi.items[n-1] = nil
	pqi.items = pqi.items[:n-1]
	return *item
}

// General compare functions
func CompareFuncIntAsc(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func CompareFuncIntDesc(i, j interface{}) bool {
	return i.(int) > j.(int)
}

func CompareFuncStringAsc(i, j interface{}) bool {
	return i.(string) < j.(string)
}

func CompareFuncStringDesc(i, j interface{}) bool {
	return i.(string) > j.(string)
}