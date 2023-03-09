package datastructure

// Queue represents a FIFO queue
type Queue []int

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(n int) {
	*q = append(*q, n)
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	n := (*q)[0]
	*q = (*q)[1:]
	return n, true
}
