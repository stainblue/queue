package queue

import (
	"testing"
)
func TestPriorityQueueIntAsc(t *testing.T) {
	testcase := struct {
		input []int
		want  []int
	}{
		input: []int{3, 6, 1, 2, 11},
		want:  []int{1, 2, 3, 6, 11},
	}

	pq := NewPriorityQueue(CompareFuncIntAsc)
	for _, v := range testcase.input {
		pq.Push(v)
	}

	for i := 0; !pq.IsEmpty(); i++ {
		if i >= len(testcase.want) {
			t.Errorf("index : %d => Pop() = %d, want null", i, pq.Pop().(int))
		} else if got, want := pq.Pop().(int), testcase.want[i]; got != want {
			t.Errorf("index : %d => Pop() = %d, want %d", i, got, want)
		}
	}
}

func TestPriorityQueueIntDesc(t *testing.T) {
	testcase := struct {
		input []int
		want  []int
	}{
		input: []int{3, 6, 1, 2, 11},
		want:  []int{11, 6, 3, 2, 1},
	}

	pq := NewPriorityQueue(CompareFuncIntDesc)
	for _, v := range testcase.input {
		pq.Push(v)
	}

	for i := 0; !pq.IsEmpty(); i++ {
		if i >= len(testcase.want) {
			t.Errorf("index : %d => Pop() = %d, want null", i, pq.Pop().(int))
		} else if got, want := pq.Pop().(int), testcase.want[i]; got != want {
			t.Errorf("index : %d => Pop() = %d, want %d", i, got, want)
		}
	}
}
