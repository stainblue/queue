package queue

import "testing"

func TestQueue(t *testing.T) {
	testcase := struct {
		input []int
		want  []int
	}{
		input: []int{3, 6, 1, 2, 11},
		want:  []int{3, 6, 1, 2, 11},
	}

	q := New()
	for _, v := range testcase.input {
		q.Enqueue(v)
	}

	for i := 0; !q.IsEmpty(); i++ {
		if i >= len(testcase.want) {
			t.Errorf("index : %d => Dequeue() = %d, want null", i, q.Dequeue().(int))
		} else if got, want := q.Dequeue().(int), testcase.want[i]; got != want {
			t.Errorf("index : %d => Dequeue() = %d, want = %d", i, got, want)
		}
	}
}